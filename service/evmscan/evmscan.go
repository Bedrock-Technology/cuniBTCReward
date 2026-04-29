package evmscan

import (
	"context"
	"cuniBTCReward/clientrpc/goeth"
	"cuniBTCReward/model"
	"cuniBTCReward/pkg/gormz"
	"cuniBTCReward/pkg/slack"
	"cuniBTCReward/service/contract/airdrop"
	"cuniBTCReward/service/contract/cunibtcvault"
	"cuniBTCReward/service/contract/delayredeemrouter"
	"cuniBTCReward/service/contract/factory"
	"cuniBTCReward/service/evmscan/config"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Scanner struct {
	database        *gorm.DB
	config          *config.EvmScanConf
	evmClients      []*EvmClient
	cuniBTCVaultAbi *abi.ABI
	redeemRouterAbi *abi.ABI
	airDropAbi      *abi.ABI
	factoryAbi      *abi.ABI
}

type EvmClient struct {
	Client  *ethclient.Client
	Factory *factory.Factory
}

func NewScanner(c *config.EvmScanConf) *Scanner {
	gormConfig := &gorm.Config{}
	if c.SqlLog {
		gormConfig.Logger = gormz.NewGormLogger()
	}
	db, err := gorm.Open(mysql.Open(c.DataSource), gormConfig)
	logx.Must(err)

	evmClients := lo.Map(c.ChainInfo, func(item config.ChainInfo, index int) *EvmClient {
		client := goeth.NewClient(item.Client.Host, item.Client.Request, item.Client.PeriodSec)
		factory, err := factory.NewFactory(common.HexToAddress(item.Factory), client)
		logx.Must(err)
		return &EvmClient{
			Client:  client,
			Factory: factory,
		}
	})

	cuniBTCVaultAbi, _ := cunibtcvault.CunibtcvaultMetaData.GetAbi()
	redeemRouterAbi, _ := delayredeemrouter.DelayredeemrouterMetaData.GetAbi()
	airDropAbi, _ := airdrop.AirdropMetaData.GetAbi()
	factoryAbi, _ := factory.FactoryMetaData.GetAbi()

	return &Scanner{
		database:        db,
		config:          c,
		evmClients:      evmClients,
		cuniBTCVaultAbi: cuniBTCVaultAbi,
		redeemRouterAbi: redeemRouterAbi,
		airDropAbi:      airDropAbi,
		factoryAbi:      factoryAbi,
	}
}

func (s *Scanner) LogScan() {
	for k, chain := range s.config.ChainInfo {
		logx.Infof("chain: %v, name: %v, logscan", chain.Client.ChainId, chain.Client.ChainName)

		cursor, err := model.GetCursor(s.database, chain.Client.ChainId)
		if err != nil {
			logx.Errorf("get chain: %v, cursor failed, err: %v", chain.Client.ChainId, err)
			return
		}

		start, end, err := s.getScanRange(s.evmClients[k].Client, cursor.BlockNumber)
		if err != nil {
			logx.Errorf("get chain: %v, latest block number failed, err: %v", chain.Client.ChainId, err)
			continue
		}
		if start == 0 || end == 0 {
			logx.Infof("chain: %v, no new block", chain.Client.ChainId)
			continue
		}
		logx.Infof("chain: %v, need scan blocks start:%v, end:%v", chain.Client.ChainId, start, end)

		factoryLogs, err := s.fetchLogs(s.evmClients[k].Client, start, end, []common.Address{common.HexToAddress(chain.Factory)})
		if err != nil {
			logx.Errorf("get chain: %v, filter logs failed, err: %v", chain.Client.ChainId, err)
			continue
		}
		logx.Infof("chain: %v, fetched factoryLogs: %v", chain.Client.ChainId, len(factoryLogs))
		// Process Factory events
		if err := s.processFactoryLog(factoryLogs, chain, s.evmClients[k]); err != nil {
			logx.Errorf("processFactoryLog error, err: %v", err)
			continue
		}

		stratigies, err := model.GetStrategy(s.database, chain.Client.ChainId)
		if err != nil {
			logx.Errorf("get chain: %v, strategy failed, err: %v", chain.Client.ChainId, err)
			continue
		}
		addresses := lo.FlatMap(stratigies, func(item model.Strategy, index int) []common.Address {
			return []common.Address{
				common.HexToAddress(item.Vault),
				common.HexToAddress(item.DelayRedeemRouter),
				common.HexToAddress(item.Airdrop),
			}
		})
		logx.Infof("chain: %v, need scan addresses: %v", chain.Client.ChainId, addresses)
		logs, err := s.fetchLogs(s.evmClients[k].Client, start, end, addresses)
		if err != nil {
			logx.Errorf("get chain: %v, filter logs failed, err: %v", chain.Client.ChainId, err)
			continue
		}
		logx.Infof("chain: %v, fetched logs: %v", chain.Client.ChainId, len(logs))
		cursor.BlockNumber = uint64(end)
		if err := s.processAndSave(cursor, logs, s.evmClients[k], chain, stratigies); err != nil {
			logx.Errorf("chain: %v, process logs failed, err: %v", chain.Client.ChainId, err)
			continue
		}
		if err := s.EpochSpin(s.evmClients[k], chain, cursor.BlockNumber, stratigies); err != nil {
			logx.Errorf("chain: %v, epoch spin failed, err: %v", chain.Client.ChainId, err)
			continue
		}
	}
}

func (s *Scanner) EpochSpin(evmClient *EvmClient, chainInfo config.ChainInfo, blockNumber uint64, stratigies []model.Strategy) error {
	for _, strategy := range stratigies {
		var epoch []model.Epoch
		err := s.database.Where("chain_id = ? AND contract = ?", chainInfo.Client.ChainId, strategy.Vault).
			Order("epoch desc").Limit(1).Find(&epoch).Error
		if err != nil {
			logx.Errorf("get epoch failed, err: %v", err)
			return err
		}
		cuniBTCVault, _ := cunibtcvault.NewCunibtcvault(common.HexToAddress(strategy.Vault), evmClient.Client)
		startGenesis, err := cuniBTCVault.StartGenesis(&bind.CallOpts{Context: context.Background()})
		if err != nil {
			logx.Errorf("get start genesis failed, err: %v", err)
			return err
		}
		operatePeriod, err := cuniBTCVault.OperatePeriod(&bind.CallOpts{Context: context.Background()})
		if err != nil {
			logx.Errorf("get operate period failed, err: %v", err)
			return err
		}
		lockupPeriod, err := cuniBTCVault.LockupPeriod(&bind.CallOpts{Context: context.Background()})
		if err != nil {
			logx.Errorf("get lockup period failed, err: %v", err)
			return err
		}
		epochNumber := (blockNumber - startGenesis.Uint64()) / (operatePeriod.Uint64() + lockupPeriod.Uint64())

		if len(epoch) == 0 {
			for i := uint64(0); i <= epochNumber; i++ {
				newEpoch := model.Epoch{
					ChainId:       chainInfo.Client.ChainId,
					Epoch:         i,
					Contract:      strategy.Vault,
					OperateStart:  startGenesis.Uint64() + i*(operatePeriod.Uint64()+lockupPeriod.Uint64()),
					LockupStart:   startGenesis.Uint64() + i*(operatePeriod.Uint64()+lockupPeriod.Uint64()) + operatePeriod.Uint64(),
					StartGenesis:  startGenesis.Uint64(),
					OperatePeriod: operatePeriod.Uint64(),
					LockupPeriod:  lockupPeriod.Uint64(),
				}
				if err := s.database.Create(&newEpoch).Error; err != nil {
					logx.Errorf("create epoch failed, err: %v", err)
					return err
				}
				slack.SendTo(s.config.NotifySlack, fmt.Sprintf("[%s] New epoch created for vault: %s, epoch: %d, operate start: %d, lockup start: %d", s.config.Name,
					strategy.Vault, newEpoch.Epoch, newEpoch.OperateStart, newEpoch.LockupStart))
			}
		} else {
			//check
			if epoch[0].OperatePeriod != operatePeriod.Uint64() || epoch[0].LockupPeriod != lockupPeriod.Uint64() || epoch[0].StartGenesis != startGenesis.Uint64() {
				return fmt.Errorf("epoch parameters updated for vault: %s", strategy.Vault)
			}
			for i := uint64(epoch[0].Epoch + 1); i <= epochNumber; i++ {
				newEpoch := model.Epoch{
					ChainId:       chainInfo.Client.ChainId,
					Epoch:         i,
					Contract:      strategy.Vault,
					OperateStart:  startGenesis.Uint64() + i*(operatePeriod.Uint64()+lockupPeriod.Uint64()),
					LockupStart:   startGenesis.Uint64() + i*(operatePeriod.Uint64()+lockupPeriod.Uint64()) + operatePeriod.Uint64(),
					StartGenesis:  startGenesis.Uint64(),
					OperatePeriod: operatePeriod.Uint64(),
					LockupPeriod:  lockupPeriod.Uint64(),
				}
				if err := s.database.Create(&newEpoch).Error; err != nil {
					logx.Errorf("create epoch failed, err: %v", err)
					return err
				}
				slack.SendTo(s.config.NotifySlack, fmt.Sprintf("[%s] New epoch created for vault: %s, epoch: %d, operate start: %d, lockup start: %d", s.config.Name,
					strategy.Vault, newEpoch.Epoch, newEpoch.OperateStart, newEpoch.LockupStart))
			}
		}
	}
	_ = s.updateEpochTimestamp(evmClient, chainInfo, blockNumber)
	return nil
}

func (s *Scanner) updateEpochTimestamp(evmClient *EvmClient, chainInfo config.ChainInfo, blockNumber uint64) error {
	//update startGenesis timestamp
	var epochStartGenesisEmpty []model.Epoch
	err := s.database.Where("chain_id = ?", chainInfo.Client.ChainId).Where("start_genesis <= ?", blockNumber).
		Where("start_genesis_timestamp = ?", 0).Find(&epochStartGenesisEmpty).Error
	if err != nil {
		return err
	}
	for _, v := range epochStartGenesisEmpty {
		//get block timestamp
		block, err := evmClient.Client.BlockByNumber(context.Background(), big.NewInt(int64(v.StartGenesis)))
		if err != nil {
			continue
		}
		v.StartGenesisTimestamp = block.Header().Time
		s.database.Save(v)
	}
	var epochOperateEmpty []model.Epoch
	err = s.database.Where("chain_id = ?", chainInfo.Client.ChainId).Where("operate_start <= ?", blockNumber).
		Where("operate_start_timestamp = ?", 0).Find(&epochOperateEmpty).Error
	if err != nil {
		return err
	}
	for _, v := range epochOperateEmpty {
		//get block timestamp
		block, err := evmClient.Client.BlockByNumber(context.Background(), big.NewInt(int64(v.OperateStart)))
		if err != nil {
			continue
		}
		v.OperateStartTimestamp = block.Header().Time
		s.database.Save(v)
	}
	var epochLockupEmpty []model.Epoch
	err = s.database.Where("chain_id = ?", chainInfo.Client.ChainId).Where("lockup_start <= ?", blockNumber).
		Where("lockup_start_timestamp = ?", 0).Find(&epochLockupEmpty).Error
	if err != nil {
		return err
	}
	for _, v := range epochLockupEmpty {
		//get block timestamp
		block, err := evmClient.Client.BlockByNumber(context.Background(), big.NewInt(int64(v.LockupStart)))
		if err != nil {
			continue
		}
		v.LockupStartTimestamp = block.Header().Time
		s.database.Save(v)
	}
	return nil
}

func (s *Scanner) processAndSave(cursor *model.Cursor, logs []types.Log, evmClient *EvmClient, chainInfo config.ChainInfo, strategies []model.Strategy) error {
	tx := s.database.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		logx.Errorf("chain: %v, start transaction failed, err: %v", chainInfo.Client.ChainId, err)
		return err
	}

	if err := s.processLogs(logs, evmClient, chainInfo, strategies, tx); err != nil {
		logx.Errorf("process logs failed, err: %v", err)
		tx.Rollback()
		return err
	}
	if err := tx.Save(cursor).Error; err != nil {
		logx.Errorf("save cursor failed, err: %v", err)
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		logx.Errorf("chain: %v, commit transaction failed, err: %v", chainInfo.Client.ChainId, err)
		return err
	}
	return nil
}

func (s *Scanner) processLogs(logs []types.Log, evmClient *EvmClient, chainInfo config.ChainInfo, strategies []model.Strategy, tx *gorm.DB) error {
	for _, log := range logs {
		if log.Removed {
			logx.Errorf("log removed, hash:%v, blockNumber:%v, blockHash:%v", log.TxHash, log.BlockNumber, log.BlockHash)
			return fmt.Errorf("log removed")
		}
		transactionRecipient, err := evmClient.Client.TransactionReceipt(context.Background(), log.TxHash)
		if err != nil {
			logx.Errorf("get transaction receipt failed, err: %v", err)
			return err
		}
		if transactionRecipient.Status != types.ReceiptStatusSuccessful {
			logx.Infof("transaction status not successful, hash: %v, status: %v", log.TxHash.Hex(), transactionRecipient.Status)
			continue
		}
		if isCuniBTCVault(log.Address, strategies) {
			// Process CuniBTCVault events
			err := s.processCuniBTCVaultLog(log, chainInfo, evmClient, tx)
			if err != nil {
				logx.Errorf("processCuniBTCVaultLog error, err: %v", err)
				return err
			}
		} else if isDelayRedeemRouter(log.Address, strategies) {
			// Process DelayRedeemRouter events
			err := s.processDelayRedeemRouterLog(log, chainInfo, evmClient, tx)
			if err != nil {
				logx.Errorf("processDelayRedeemRouterLog error, err: %v", err)
				return err
			}
		} else if isAirDrop(log.Address, strategies) {
			// Process Airdrop events
			err := s.processAirDropLog(log, chainInfo, evmClient, tx)
			if err != nil {
				logx.Errorf("processAirDropLog error, err: %v", err)
				return err
			}
		} else {
			continue
		}
	}
	return nil
}

func isCuniBTCVault(address common.Address, strategies []model.Strategy) bool {
	for _, strategy := range strategies {
		if common.HexToAddress(strategy.Vault) == address {
			return true
		}
	}
	return false
}

func isDelayRedeemRouter(address common.Address, strategies []model.Strategy) bool {
	for _, strategy := range strategies {
		if common.HexToAddress(strategy.DelayRedeemRouter) == address {
			return true
		}
	}
	return false
}

func isAirDrop(address common.Address, strategies []model.Strategy) bool {
	for _, strategy := range strategies {
		if common.HexToAddress(strategy.Airdrop) == address {
			return true
		}
	}
	return false
}

func (s *Scanner) processFactoryLog(logs []types.Log, chainInfo config.ChainInfo, evmClient *EvmClient) error {
	for _, log := range logs {
		if log.Removed {
			logx.Errorf("log removed, hash:%v, blockNumber:%v, blockHash:%v", log.TxHash, log.BlockNumber, log.BlockHash)
			return fmt.Errorf("log removed")
		}
		transactionRecipient, err := evmClient.Client.TransactionReceipt(context.Background(), log.TxHash)
		if err != nil {
			logx.Errorf("get transaction receipt failed, err: %v", err)
			return err
		}
		if transactionRecipient.Status != types.ReceiptStatusSuccessful {
			logx.Infof("transaction status not successful, hash: %v, status: %v", log.TxHash.Hex(), transactionRecipient.Status)
			continue
		}
		eventName, err := s.factoryAbi.EventByID(log.Topics[0])
		if err != nil {
			logx.Errorf("get event name failed, maybe upgraded hash: %v, err: %v", log.TxHash, err)
			return nil
		}
		switch eventName.Name {
		case "StrategyCreate":
			newStrategyEvent, err := evmClient.Factory.ParseStrategyCreate(log)
			if err != nil {
				logx.Errorf("parse new strategy event failed, err: %v", err)
				return err
			}
			strategy := &model.Strategy{
				ChainId:           chainInfo.Client.ChainId,
				Name:              newStrategyEvent.Strategy.Name,
				Symbol:            newStrategyEvent.Strategy.Symbol,
				CuniBTC:           newStrategyEvent.Strategy.CuniBTC.String(),
				Vault:             newStrategyEvent.Strategy.Vault.String(),
				DelayRedeemRouter: newStrategyEvent.Strategy.DelayRedeemRouter.String(),
				Airdrop:           newStrategyEvent.Strategy.Airdrop.String(),
			}
			slack.SendTo(s.config.NotifySlack, fmt.Sprintf("[%s] New strategy created: %s, symbol: %s, vault: %s", s.config.Name, strategy.Name, strategy.Symbol, strategy.Vault))
			if err := s.database.Clauses(clause.OnConflict{DoNothing: true}).Create(&strategy).Error; err != nil {
				logx.Errorf("create strategy failed, err: %v", err)
				return err
			}
		}
	}
	return nil
}

func (s *Scanner) processCuniBTCVaultLog(log types.Log, chainInfo config.ChainInfo, evmClient *EvmClient, tx *gorm.DB) error {
	eventName, err := s.cuniBTCVaultAbi.EventByID(log.Topics[0])
	if err != nil {
		logx.Errorf("get event name failed, maybe upgraded hash: %v, err: %v", log.TxHash, err)
		return nil
	}
	cuniBTCVault, _ := cunibtcvault.NewCunibtcvault(log.Address, evmClient.Client)
	switch eventName.Name {
	case "Minted":
		mintedEvent, err := cuniBTCVault.ParseMinted(log)
		if err != nil {
			logx.Errorf("parse minted event failed, err: %v", err)
			return nil
		}
		return tx.Create(&model.EvmTransaction{
			Address:        mintedEvent.Sender.String(),
			ChainId:        chainInfo.Client.ChainId,
			Hash:           log.TxHash.String(),
			IndexInHash:    log.Index,
			Contract:       log.Address.String(),
			Token:          mintedEvent.Token.String(),
			BlockNumber:    log.BlockNumber,
			BlockTimestamp: log.BlockTimestamp,
			Amount:         decimal.NewFromBigInt(mintedEvent.Amount, 0),
			LogMethod:      "Minted",
		}).Error
	case "PeriodSet":
		periodSetEvent, err := cuniBTCVault.ParsePeriodSet(log)
		if err != nil {
			logx.Errorf("parse period set event failed, err: %v", err)
			return nil
		}
		slack.SendTo(s.config.NotifySlack, fmt.Sprintf("[%s] Period updated for vault: %s, operate period: %d, lockup period: %d", s.config.Name,
			log.Address.String(), periodSetEvent.OperatePeriod, periodSetEvent.LockupPeriod))
	}
	return nil
}

func (s *Scanner) processDelayRedeemRouterLog(log types.Log, chainInfo config.ChainInfo, evmClient *EvmClient, tx *gorm.DB) error {
	eventName, err := s.redeemRouterAbi.EventByID(log.Topics[0])
	if err != nil {
		logx.Errorf("get event name failed, maybe upgraded hash: %v, err: %v", log.TxHash, err)
		return nil
	}
	redeemRouter, _ := delayredeemrouter.NewDelayredeemrouter(log.Address, evmClient.Client)
	switch eventName.Name {
	case "DelayedRedeemCreated":
		redeemCreatedEvent, err := redeemRouter.ParseDelayedRedeemCreated(log)
		if err != nil {
			logx.Errorf("parse delayed redeem created event failed, err: %v", err)
			return err
		}

		if err := tx.Create(&model.EvmTransaction{
			Address:        redeemCreatedEvent.Recipient.String(),
			ChainId:        chainInfo.Client.ChainId,
			Hash:           log.TxHash.String(),
			IndexInHash:    log.Index,
			Contract:       log.Address.String(),
			Token:          redeemCreatedEvent.Token.String(),
			BlockNumber:    log.BlockNumber,
			BlockTimestamp: log.BlockTimestamp,
			Amount: decimal.NewFromBigInt(redeemCreatedEvent.Amount, 0).
				Add(decimal.NewFromBigInt(redeemCreatedEvent.RedeemFee, 0)).
				Neg(),
			LogMethod: "DelayedRedeemCreated",
		}).Error; err != nil {
			return err
		}
		if err := tx.Create(&model.DelayRedeemRecord{
			ChainId:           chainInfo.Client.ChainId,
			Address:           redeemCreatedEvent.Recipient.String(),
			IndexInHash:       log.Index,
			Contract:          log.Address.String(),
			Token:             redeemCreatedEvent.Token.String(),
			Amount:            decimal.NewFromBigInt(redeemCreatedEvent.Amount, 0),
			Fee:               decimal.NewFromBigInt(redeemCreatedEvent.RedeemFee, 0),
			CreateHash:        log.TxHash.String(),
			CreateBlockTime:   time.Unix(int64(log.BlockTimestamp), 0),
			CreateBlockNumber: log.BlockNumber,
			Claimed:           false,
			ClaimTx:           "",
			ClaimAt:           time.Unix(0, 0),
			Index:             redeemCreatedEvent.Index.Uint64(),
		}).Error; err != nil {
			return err
		}
	case "DelayedRedeemsCompleted":
		redeemCompleteEvent, err := redeemRouter.ParseDelayedRedeemsCompleted(log)
		if err != nil {
			logx.Errorf("parse delayed redeem completed event failed, err: %v", err)
			return err
		}
		return tx.Model(&model.DelayRedeemRecord{}).Where("chain_id = ? AND contract = ? AND address = ? AND `index` < ?",
			chainInfo.Client.ChainId, log.Address.String(), redeemCompleteEvent.Recipient.String(),
			redeemCompleteEvent.DelayedRedeemsCompleted.Uint64()).Where("claimed = ?", false).
			Updates(map[string]interface{}{
				"claimed":  true,
				"claim_tx": log.TxHash.String(),
				"claim_at": time.Unix(int64(log.BlockTimestamp), 0),
			}).Error
	}
	return nil
}

func (s *Scanner) processAirDropLog(log types.Log, chainInfo config.ChainInfo, evmClient *EvmClient, tx *gorm.DB) error {
	eventName, err := s.airDropAbi.EventByID(log.Topics[0])
	if err != nil {
		logx.Errorf("get event name failed, maybe upgraded hash: %v, err: %v", log.TxHash, err)
		return nil
	}
	airDrop, _ := airdrop.NewAirdrop(log.Address, evmClient.Client)
	switch eventName.Name {
	case "AirdropClaimed":
		claimedEvent, err := airDrop.ParseAirdropClaimed(log)
		if err != nil {
			logx.Errorf("parse air drop claimed event failed, err: %v", err)
			return err
		}
		return tx.Model(&model.AirDropRecord{}).Where("chain_id = ? AND contract = ? AND epoch = ? AND address = ?",
			chainInfo.Client.ChainId, log.Address.String(), claimedEvent.Epoch.Uint64(), claimedEvent.User.String()).
			Updates(map[string]interface{}{
				"claimed":  true,
				"claim_tx": log.TxHash.String(),
				"claim_at": time.Unix(int64(log.BlockTimestamp), 0),
			}).Error
	case "MerkleRootSubmit":
		rootEvent, err := airDrop.ParseMerkleRootSubmit(log)
		if err != nil {
			logx.Errorf("parse merkle root submit event failed, err: %v", err)
			return err
		}
		slack.SendTo(s.config.NotifySlack, fmt.Sprintf("[%s] New merkle root submitted for vault: %s, epoch: %d", s.config.Name, log.Address.String(), rootEvent.Epoch.Int64()))
		return tx.Create(&model.AirDropEpoch{
			ChainId:   chainInfo.Client.ChainId,
			Contract:  log.Address.String(),
			Epoch:     uint64(rootEvent.Epoch.Int64()),
			Root:      hexutil.Encode(rootEvent.Root[:]),
			Token:     rootEvent.Token.String(),
			ValidTime: rootEvent.RewardsValidTime.Uint64(),
			ActiveAt:  time.Unix(int64(rootEvent.ActivatedAt.Uint64()), 0),
			Disabled:  false,
		}).Error
	case "MerkleRootUpdate":
		rootUpdateEvent, err := airDrop.ParseMerkleRootUpdate(log)
		if err != nil {
			logx.Errorf("parse merkle root update event failed, err: %v", err)
			return err
		}
		slack.SendTo(s.config.NotifySlack, fmt.Sprintf("[%s] Merkle root updated for vault: %s, epoch: %d", s.config.Name, log.Address.String(), rootUpdateEvent.Epoch.Int64()))
		return tx.Model(&model.AirDropEpoch{}).Where("chain_id = ? AND contract = ? AND epoch = ?",
			chainInfo.Client.ChainId, log.Address.String(), rootUpdateEvent.Epoch.Uint64()).
			Updates(map[string]interface{}{"root": hexutil.Encode(rootUpdateEvent.Root[:])}).Error
	case "TokenUpdate":
		tokenUpdateEvent, err := airDrop.ParseTokenUpdate(log)
		if err != nil {
			logx.Errorf("parse token update event failed, err: %v", err)
			return err
		}
		slack.SendTo(s.config.NotifySlack, fmt.Sprintf("[%s] Token updated for vault: %s, epoch: %d", s.config.Name, log.Address.String(), tokenUpdateEvent.Epoch.Int64()))
		return tx.Model(&model.AirDropEpoch{}).Where("chain_id = ? AND contract = ? AND epoch = ?",
			chainInfo.Client.ChainId, log.Address.String(), tokenUpdateEvent.Epoch.Uint64()).
			Updates(map[string]interface{}{"token": tokenUpdateEvent.Token.String()}).Error
	case "ValidDurationUpdate":
		durationUpdateEvent, err := airDrop.ParseValidDurationUpdate(log)
		if err != nil {
			logx.Errorf("parse valid duration update event failed, err: %v", err)
			return err
		}
		slack.SendTo(s.config.NotifySlack, fmt.Sprintf("[%s] Valid duration updated for vault: %s, epoch: %d", s.config.Name,
			log.Address.String(), durationUpdateEvent.Epoch.Int64()))
		return tx.Model(&model.AirDropEpoch{}).Where("chain_id = ? AND contract = ? AND epoch = ?",
			chainInfo.Client.ChainId, log.Address.String(), durationUpdateEvent.Epoch.Uint64()).
			Updates(map[string]interface{}{"valid_time": durationUpdateEvent.ValidDuration.Uint64()}).Error
	case "DistributionDisabledSet":
		disabledEvent, err := airDrop.ParseDistributionDisabledSet(log)
		if err != nil {
			logx.Errorf("parse distribution disabled set event failed, err: %v", err)
			return err
		}
		slack.SendTo(s.config.NotifySlack, fmt.Sprintf("[%s] Distribution disabled set for vault: %s, epoch: %d, status: %t",
			s.config.Name, log.Address.String(), disabledEvent.Epoch.Int64(), disabledEvent.Status))
		return tx.Model(&model.AirDropEpoch{}).Where("chain_id = ? AND contract = ? AND epoch = ?",
			chainInfo.Client.ChainId, log.Address.String(), disabledEvent.Epoch.Uint64()).
			Updates(map[string]interface{}{"disabled": disabledEvent.Status}).Error
	}
	return nil
}

func (s *Scanner) getScanRange(client *ethclient.Client, cursorBlock uint64) (int64, int64, error) {
	latestBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, 0, err
	}
	latestBlockNumber = latestBlockNumber - 1 //delay block
	needScanBlocksNumber := int64(latestBlockNumber) - int64(cursorBlock)
	if needScanBlocksNumber <= 0 {
		return 0, 0, nil
	}
	if needScanBlocksNumber > 1000 {
		needScanBlocksNumber = 1000
	}
	start := int64(cursorBlock + 1)
	end := int64(cursorBlock) + needScanBlocksNumber
	return start, end, nil
}

func (s *Scanner) fetchLogs(client *ethclient.Client, start, end int64, addresses []common.Address) ([]types.Log, error) {
	return client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(start),
		ToBlock:   big.NewInt(end),
		Addresses: addresses,
	})
}

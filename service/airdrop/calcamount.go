package airdrop

import (
	"context"
	"cuniBTCReward/clientrpc/goeth"
	"cuniBTCReward/model"
	"cuniBTCReward/pkg/gormz"
	"cuniBTCReward/service/airdrop/config"
	"fmt"
	"math/big"

	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CalcAvgAmount startAmount is the amount < startBlock
func CalcAvgAmount(startAmount decimal.Decimal, transactions []*model.EvmTransaction, startBlock, endBlock uint64) (decimal.Decimal, error) {
	if startBlock >= endBlock {
		return decimal.Zero, fmt.Errorf("startBlock[%d] >= endBlock[%d]", startBlock, endBlock)
	}
	if startAmount.IsNegative() {
		return decimal.Zero, fmt.Errorf("startAmount[%s] is negative", startAmount.String())
	}
	if len(transactions) == 0 {
		return startAmount, nil
	}
	totalBlocks := endBlock - startBlock
	var weightedSum decimal.Decimal
	var previousAmount = startAmount
	var previousBlock = startBlock

	for _, tx := range transactions {
		if tx.BlockNumber < startBlock || tx.BlockNumber > endBlock {
			continue
		}
		blockDelta := tx.BlockNumber - previousBlock
		weightedSum = weightedSum.Add(previousAmount.Mul(decimal.NewFromInt(int64(blockDelta))))
		previousAmount = previousAmount.Add(tx.Amount)
		previousBlock = tx.BlockNumber
		if previousAmount.IsNegative() {
			return decimal.Zero, fmt.Errorf("previousAmount[%s] is negative, txHash[%s]", previousAmount.String(), tx.Hash)
		}
	}

	blockDelta := endBlock - previousBlock
	weightedSum = weightedSum.Add(previousAmount.Mul(decimal.NewFromInt(int64(blockDelta))))

	return weightedSum.Div(decimal.NewFromInt(int64(totalBlocks))).Floor(), nil
}

type Airdrop struct {
	database   *gorm.DB
	config     *config.AirdropConf
	evmClients map[uint]*EvmClient
}
type EvmClient struct {
	Client *ethclient.Client
}

func NewAirdrop(c *config.AirdropConf) *Airdrop {
	gormConfig := &gorm.Config{}
	if c.SqlLog {
		gormConfig.Logger = gormz.NewGormLogger()
	}
	db, err := gorm.Open(mysql.Open(c.DataSource), gormConfig)
	logx.Must(err)

	return &Airdrop{
		database: db,
		config:   c,
		evmClients: lo.SliceToMap(c.ChainInfo, func(item config.ChainInfo) (uint, *EvmClient) {
			client := goeth.NewClient(item.Client.Host, item.Client.Request, item.Client.PeriodSec)
			return item.Client.ChainId, &EvmClient{
				Client: client,
			}
		}),
	}
}

type AvgAmountResult struct {
	Address   string
	AvgAmount decimal.Decimal
}

func (a *Airdrop) GetCursor(chainid uint) (*model.Cursor, error) {
	var cursor model.Cursor
	err := a.database.Model(&model.Cursor{}).Where("chain_id = ?", chainid).First(&cursor).Error
	return &cursor, err
}

func (a *Airdrop) GetAvgAmount(startBlock, endBlock uint64, chainId uint, contract []string, token string) (result []AvgAmountResult, err error) {
	startAmount, err := a.getAllAddressStartAmount(startBlock, chainId, contract, token)
	if err != nil {
		return nil, err
	}
	for _, item := range startAmount {
		if item.AvgAmount.IsNegative() {
			return nil, fmt.Errorf("startAmount[%s] is negative for address[%s]", item.AvgAmount.String(), item.Address)
		}
		transactions, err := a.getAddressTransactions(item.Address, startBlock, endBlock, chainId)
		if err != nil {
			return nil, err
		}
		avgAmount, err := CalcAvgAmount(item.AvgAmount, transactions, startBlock, endBlock)
		if err != nil {
			return nil, fmt.Errorf("calc avg amount failed for address[%s], err: %v", item.Address, err)
		}
		result = append(result, AvgAmountResult{
			Address:   item.Address,
			AvgAmount: avgAmount,
		})
	}
	return
}

func (a *Airdrop) getAllAddressStartAmount(startBlock uint64, chainId uint, contract []string, token string) (result []AvgAmountResult, err error) {
	var distinctAddresses []string
	ret := a.database.Model(&model.EvmTransaction{}).Distinct("address").
		Where("chain_id = ?", chainId).Pluck("address", &distinctAddresses)
	if ret.Error != nil {
		return nil, ret.Error
	}

	var addressAmounts []AvgAmountResult
	ret = a.database.Model(&model.EvmTransaction{}).Select("address, sum(amount) as avg_amount").
		Where("chain_id = ? and block_number < ? and contract in (?) and token = ?", chainId, startBlock, contract, token).
		Group("address").Order("address").Scan(&addressAmounts)
	if ret.Error != nil {
		return nil, ret.Error
	}

	addressAmountMap := lo.SliceToMap(addressAmounts, func(item AvgAmountResult) (string, AvgAmountResult) {
		return item.Address, item
	})

	result = lo.Map(distinctAddresses, func(address string, index int) AvgAmountResult {
		if item, ok := addressAmountMap[address]; ok {
			return item
		}
		return AvgAmountResult{
			Address:   address,
			AvgAmount: decimal.Zero,
		}
	})
	return
}

func (c *Airdrop) getAddressTransactions(address string, startBlock, endBlock uint64, chainId uint) (result []*model.EvmTransaction, err error) {
	ret := c.database.Model(&model.EvmTransaction{}).
		Where("chain_id = ? and address = ? and block_number >= ? and block_number <= ?", chainId, address, startBlock, endBlock).
		Order("block_number").
		Find(&result)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return
}

func (c *Airdrop) GetAllAddressAtBlock(blockNumber uint64, chainId uint, balanceCheck bool, contract []string, token string) (result []AvgAmountResult, err error) {
	//sum amount for each address where block_number <= blockNumber, group by address
	ret := c.database.Model(&model.EvmTransaction{}).Select("address, sum(amount) as avg_amount").
		Where("chain_id = ? and block_number <= ? and contract in (?) and token = ?", chainId, blockNumber, contract, token).
		Group("address").Order("address").Scan(&result)
	if ret.Error != nil {
		return nil, ret.Error
	}

	if balanceCheck {
		for _, item := range result {
			balance, err := c.GetTokenBalanceAtBlock(item.Address, blockNumber, chainId, token)
			if err != nil {
				return nil, fmt.Errorf("get token balance at block failed for address[%s], err: %v", item.Address, err)
			}
			if !balance.Equal(item.AvgAmount) {
				return nil, fmt.Errorf("balance check failed for address[%s]", item.Address)
			}
			// logx.Infof("balance check passed for address[%s], balance: %s, amount: %s", item.Address, balance.String(), item.AvgAmount.String())
		}
	}
	return
}

func (c *Airdrop) GetTokenBalanceAtBlock(address string, blockNumber uint64, chainId uint, token string) (decimal.Decimal, error) {
	evmClient, ok := c.evmClients[chainId]
	if !ok {
		return decimal.Zero, fmt.Errorf("evm client not found for chainId[%d]", chainId)
	}
	tokenAddress := common.HexToAddress(token)    // ERC20 Contract
	walletAddress := common.HexToAddress(address) // User Wallet

	// 1. Prepare balanceOf(address) signature
	transferFnSignature := []byte("balanceOf(address)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4] // 0x70a08231

	// 2. Pad address to 32 bytes
	paddedAddress := common.LeftPadBytes(walletAddress.Bytes(), 32)

	// 3. Combine method ID and address
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)

	// 4. Create and execute call
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	result, err := evmClient.Client.CallContract(context.Background(), msg, big.NewInt(0).SetUint64(blockNumber))
	if err != nil {
		return decimal.Zero, fmt.Errorf("call contract failed, err: %v", err)
	}

	// 5. Decode result
	balance := new(big.Int).SetBytes(result)
	return decimal.NewFromBigInt(balance, 0), nil
}

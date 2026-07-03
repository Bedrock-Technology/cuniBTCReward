// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"
	"cuniBTCReward/model"
	"cuniBTCReward/service/airdrop"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DistributeRewardsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDistributeRewardsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DistributeRewardsLogic {
	return &DistributeRewardsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type EpochRecord struct {
	Address string          `gorm:"column:address"`
	Share   decimal.Decimal `gorm:"column:share"`
	Queue   decimal.Decimal `gorm:"column:queue"`
	Amount  decimal.Decimal `gorm:"column:amount"`
}

func (l *DistributeRewardsLogic) DistributeRewards(req *types.DistributeRewardsReq) (resp *types.DistributeRewardsResp, err error) {
	chainID := uint(l.svcCtx.Config.DefaultChainId)
	//check range
	totalRewards, err := decimal.NewFromString(req.TotalRewards)
	if err != nil {
		return nil, err
	}
	totalRewards = totalRewards.Mul(decimal.New(1, 8))
	if totalRewards.IsZero() {
		return nil, fmt.Errorf("total rewards is zero for epoch %d, symbol: %s", req.Epoch, req.Symbol)
	}
	//0.1
	if totalRewards.Cmp(decimal.New(1, 7)) >= 0 {
		return nil, fmt.Errorf("totalReward too large")
	}
	// Step 1: find Strategy by symbol, get vault and airdrop contract addresses
	var strategy model.Strategy
	err = l.svcCtx.Database.WithContext(l.ctx).
		Where("chain_id = ? AND symbol = ? AND deleted_at IS NULL", chainID, req.Symbol).
		First(&strategy).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("strategy not found for symbol: %s", req.Symbol)
		}
		return nil, err
	}

	// Step 2: find Epoch by vault contract and epoch number, validate it exists and has ended
	var epoch model.Epoch
	err = l.svcCtx.Database.WithContext(l.ctx).
		Where("chain_id = ? AND contract = ? AND epoch = ? AND deleted_at IS NULL",
			chainID, strategy.Vault, req.Epoch).
		First(&epoch).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("epoch %d not found for symbol: %s", req.Epoch, req.Symbol)
		}
		return nil, err
	}

	// Check if epoch has ended: lockup_start + lockup_period < now
	epochEndTime := epoch.LockupStart + epoch.LockupPeriod
	if uint64(time.Now().Unix()) < epochEndTime {
		return nil, fmt.Errorf("epoch %d has not ended yet, end time: %d, now: %d",
			req.Epoch, epochEndTime, time.Now().Unix())
	}

	l.Infof("Processing distribution for symbol: %s, epoch: %d, vault: %s, airdrop: %s",
		req.Symbol, req.Epoch, strategy.Vault, strategy.Airdrop)

	// Step 3: query evm_transactions for vault contract, aggregate by address
	// Deposits before lockup_start = share, deposits after lockup_start = queued
	var records []EpochRecord
	sql := `SELECT t.address,
       COALESCE(SUM(CASE WHEN t.block_timestamp < ? THEN t.amount ELSE 0 END), 0) AS share,
       COALESCE(SUM(CASE WHEN t.block_timestamp >= ? AND t.block_timestamp < ? THEN t.amount ELSE 0 END), 0) AS queue
FROM evm_transactions t
WHERE t.chain_id = ? AND t.contract = ? AND t.deleted_at IS NULL AND t.amount > 0
GROUP BY t.address
HAVING share > 0 OR queue > 0`
	args := []interface{}{
		epoch.LockupStart, epoch.LockupStart, epoch.LockupStart + epoch.LockupPeriod,
		chainID, strategy.Vault,
	}
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&records).Error
	if err != nil {
		return nil, fmt.Errorf("query evm_transactions failed: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("no qualifying transactions found for epoch %d, symbol: %s", req.Epoch, req.Symbol)
	}

	// Step 4: calculate each address's amount proportionally
	totalShare := decimal.Zero
	for _, r := range records {
		totalShare = totalShare.Add(r.Share)
	}
	if totalShare.IsZero() {
		return nil, fmt.Errorf("total share is zero for epoch %d, symbol: %s", req.Epoch, req.Symbol)
	}

	for i := range records {
		records[i].Amount = records[i].Share.Div(totalShare).Mul(totalRewards)
	}
	// Get Proof
	leaves := lo.Map(records, func(item EpochRecord, _ int) airdrop.TreeLeaf {
		return airdrop.TreeLeaf{
			Address: item.Address,
			Amount:  item.Amount.String(),
		}
	})
	shares := lo.Map(records, func(item EpochRecord, _ int) string {
		return item.Share.String()
	})
	airDropRecord, merkleRoot, err := airdrop.GetAirdrop(chainID, strategy.Airdrop, req.Epoch, leaves, shares)
	if err != nil {
		return nil, fmt.Errorf("generate airdrop proof error: %v", err)
	}
	apr := totalRewards.Div(totalShare).Div(decimal.NewFromUint64(epoch.LockupPeriod)).
		Mul(decimal.NewFromInt32(31536000))
	// Step 5: upsert into air_drop_records
	for i, r := range records {
		(*airDropRecord)[i].Queued = r.Queue
	}
	err = l.svcCtx.Database.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(airDropRecord, 500).Error; err != nil {
			return err
		}
		if err := tx.Create(&model.AirDropEpoch{
			ChainId:    chainID,
			Contract:   strategy.Airdrop,
			Epoch:      req.Epoch,
			Apy:        apr.InexactFloat64(),
			Token:      l.svcCtx.Config.UniBTC,
			Root:       "",
			MerkleRoot: hexutil.Encode(merkleRoot),
			ValidTime:  0,
			ActiveAt:   time.Time{},
			Disabled:   false,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("create error: %v", err)
	}

	l.Infof("Distribution complete for symbol: %s, epoch: %d, total addresses: %d, total share: %s, total rewards: %s",
		req.Symbol, req.Epoch, len(records), totalShare.String(), totalRewards.String())

	resp = &types.DistributeRewardsResp{
		DistributeRewardsReq: *req,
		Participants:         int64(len(records)),
	}
	return
}

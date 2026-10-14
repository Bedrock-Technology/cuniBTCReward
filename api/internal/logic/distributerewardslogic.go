// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cuniBTCReward/api/internal/config"
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
	Address         string          `gorm:"column:address"`
	Share           decimal.Decimal `gorm:"column:share"`
	Queue           decimal.Decimal `gorm:"column:queue"`
	Amount          decimal.Decimal `gorm:"column:amount"`
	LastDepositTime int64           `gorm:"column:last_deposit_time"`
}

type EpochRecordExtra struct {
	Address         string          `gorm:"column:address"`
	BaseShare       decimal.Decimal `gorm:"column:base_share"`
	SharesExtra     decimal.Decimal `gorm:"column:total_extra_share"`
	AmountBase      decimal.Decimal `gorm:"column:base_profit"`
	AmountExtra     decimal.Decimal `gorm:"column:total_extra_profit"`
	AmountTotal     decimal.Decimal `gorm:"column:total_profit"`
	LastDepositTime int64           `gorm:"column:last_deposit_time"`
}

func (l *DistributeRewardsLogic) DistributeRewardsWithLock(req *types.DistributeRewardsReq) (resp *types.DistributeRewardsResp, err error) {
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
	//1
	if totalRewards.Cmp(decimal.New(1, 8)) > 0 {
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
       COALESCE(SUM(CASE WHEN t.block_timestamp >= ? AND t.block_timestamp < ? THEN t.amount ELSE 0 END), 0) AS queue,
	   MAX(CASE WHEN t.contract = ? THEN t.block_timestamp END) AS last_deposit_time
FROM evm_transactions t
WHERE t.chain_id = ? AND t.contract IN ? AND t.deleted_at IS NULL
GROUP BY t.address
HAVING share > 0 OR queue > 0`
	args := []interface{}{
		epoch.LockupStart, epoch.LockupStart, epoch.LockupStart + epoch.LockupPeriod,
		strategy.Vault,
		chainID, []string{strategy.Vault, strategy.DelayRedeemRouter},
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
		records[i].Amount = records[i].Share.Div(totalShare).Mul(totalRewards).Floor()
	}
	// Get Proof
	leaves := lo.Map(records, func(item EpochRecord, _ int) airdrop.TreeLeaf {
		return airdrop.TreeLeaf{
			Address: item.Address,
			Amount:  item.Amount.String(),
		}
	})
	airDropRecord, merkleRoot, err := airdrop.GetAirdrop(chainID, strategy.Airdrop, req.Epoch, leaves)
	if err != nil {
		return nil, fmt.Errorf("generate airdrop proof error: %v", err)
	}
	apr := totalRewards.Div(totalShare).Div(decimal.NewFromUint64(epoch.LockupPeriod)).
		Mul(decimal.NewFromInt32(31536000))
	// Step 5: upsert into air_drop_records
	for i, r := range records {
		(*airDropRecord)[i].Queued = r.Queue
		(*airDropRecord)[i].LastDepositTime = r.LastDepositTime
		(*airDropRecord)[i].AmountBase = r.Amount
		(*airDropRecord)[i].Shares = r.Share
	}
	email, _ := l.ctx.Value("email").(string)
	if email == "" {
		return nil, fmt.Errorf("email empty")
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
			ActiveAt:   time.Now().UTC().AddDate(50, 0, 0), //50 years later
			Disabled:   false,
			SubmitBy:   email,
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

func (l *DistributeRewardsLogic) DistributeRewards(req *types.DistributeRewardsReq) (resp *types.DistributeRewardsResp, err error) {
	if req.Symbol == "b2uniBTC" {
		resp, err = l.DistributeRewardsNoLock(req)
	} else {
		resp, err = l.DistributeRewardsWithLock(req)
	}
	return
}

func (l *DistributeRewardsLogic) DistributeRewardsNoLock(req *types.DistributeRewardsReq) (resp *types.DistributeRewardsResp, err error) {
	chainID := uint(l.svcCtx.Config.DefaultChainId)
	symbolStep, ok := lo.Find(l.svcCtx.Config.ExtraRewards, func(i config.ExtraReward) bool {
		return i.Symbol == req.Symbol
	})
	if !ok {
		return nil, fmt.Errorf("not found step")
	}
	//check range
	totalRewards, err := decimal.NewFromString(req.TotalRewards)
	if err != nil {
		return nil, err
	}
	totalRewards = totalRewards.Mul(decimal.New(1, 8))
	// if totalRewards.IsZero() {
	// 	return nil, fmt.Errorf("total rewards is zero for epoch %d, symbol: %s", req.Epoch, req.Symbol)
	// }
	//1
	if totalRewards.Cmp(decimal.New(1, 8)) > 0 {
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
	var records []EpochRecordExtra
	sql := `WITH user_running_balance AS (
		SELECT
        address,
        block_timestamp,
        contract,
        SUM(amount) OVER (PARTITION BY address ORDER BY block_timestamp) AS balance
    FROM evm_transactions
    WHERE block_timestamp <= ? AND contract IN ?  AND chain_id = ? AND deleted_at IS NULL -- period end; vault and delayrouter; chainid
	),
	timed_balance AS (
    SELECT
        address,
        GREATEST(block_timestamp, ?) AS start_t, -- peroid start
        LEAST(
            COALESCE(
                LEAD(block_timestamp) OVER (PARTITION BY address ORDER BY block_timestamp),
                ? -- peroid end
            ),
            ? -- peroid end
        ) AS end_t,
        balance,
        block_timestamp,
        contract
    FROM user_running_balance
),
	valid_intervals AS (
    SELECT
        address,
        balance,
        block_timestamp,
        contract,
        FLOOR((end_t - start_t) / ?) AS duration_days -- step in sec
    FROM timed_balance
    WHERE start_t < end_t AND balance > 0
),
	base_profits AS (
    SELECT
        address,
        SUM(balance * duration_days) AS user_share_sum,
        MAX(CASE WHEN contract = ? THEN block_timestamp END) AS last_deposit_time
    FROM valid_intervals
    GROUP BY address
),
	global_base AS (
    SELECT SUM(user_share_sum) AS grand_total FROM base_profits
),
	user_base_final AS (
    SELECT
        b.address,
        FLOOR(b.user_share_sum * ? / g.grand_total) AS base_profit, -- totalReward
        b.user_share_sum,
        last_deposit_time
    FROM base_profits b
    CROSS JOIN global_base g
),
	user_extra_aggregated AS (
    SELECT
        address,
        SUM(amount) AS total_extra_profit,
        SUM(share) AS total_extra_share,
        last_deposit_time
    FROM extra_reward_dis
    WHERE timestamp >= ? -- period start
      AND timestamp < ? -- peroid end
      AND symbol = ? -- symbol
    GROUP BY address
)

SELECT
    COALESCE(b.address, e.address) AS address,
    COALESCE(b.base_profit, 0) AS base_profit,
    COALESCE(e.total_extra_profit, 0) AS total_extra_profit,
    (COALESCE(b.base_profit, 0) + COALESCE(e.total_extra_profit, 0)) AS total_profit,
    b.user_share_sum AS base_share, e.total_extra_share, b.last_deposit_time
FROM user_base_final b
LEFT JOIN user_extra_aggregated e ON b.address = e.address

UNION

SELECT
    e.address,
    0 AS base_profit,
    e.total_extra_profit,
    e.total_extra_profit AS total_profit,
    0 AS base_share,
    e.total_extra_share,
    e.last_deposit_time
FROM user_extra_aggregated e
WHERE e.address NOT IN (SELECT address FROM user_base_final)
ORDER BY last_deposit_time DESC`
	args := []interface{}{
		epoch.LockupStart + epoch.LockupPeriod, []string{strategy.Vault, strategy.DelayRedeemRouter}, chainID,
		epoch.OperatePeriod, epoch.LockupStart + epoch.LockupPeriod, epoch.LockupStart + epoch.LockupPeriod,
		symbolStep,
		strategy.Vault,
		totalRewards,
		epoch.OperatePeriod, epoch.LockupStart + epoch.LockupPeriod,
		req.Symbol,
	}
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&records).Error
	if err != nil {
		return nil, fmt.Errorf("query evm_transactions failed: %w", err)
	}
	totalShare := decimal.Zero
	for _, r := range records {
		totalShare = totalShare.Add(r.AmountTotal)
	}
	if totalShare.IsZero() {
		return nil, fmt.Errorf("total share is zero for epoch %d, symbol: %s", req.Epoch, req.Symbol)
	}

	// Get Proof
	leaves := lo.Map(records, func(item EpochRecordExtra, _ int) airdrop.TreeLeaf {
		return airdrop.TreeLeaf{
			Address: item.Address,
			Amount:  item.AmountTotal.String(),
		}
	})
	airDropRecord, merkleRoot, err := airdrop.GetAirdrop(chainID, strategy.Airdrop, req.Epoch, leaves)
	if err != nil {
		return nil, fmt.Errorf("generate airdrop proof error: %v", err)
	}

	totalUserAverage := decimal.Zero
	for _, r := range records {
		totalUserAverage = totalUserAverage.Add(r.BaseShare)
	}
	shareInsec:=totalUserAverage.Div(decimal.NewFromUint64(epoch.OperatePeriod))
	apr:=
	// Step 5: upsert into air_drop_records
	for i, r := range records {
		(*airDropRecord)[i].LastDepositTime = r.LastDepositTime
		(*airDropRecord)[i].AmountBase = r.AmountBase
		(*airDropRecord)[i].Shares = r.BaseShare
		(*airDropRecord)[i].Amount = r.AmountTotal
		(*airDropRecord)[i].AmountExtra = r.AmountExtra
		(*airDropRecord)[i].SharesExtra = r.SharesExtra
	}
	email, _ := l.ctx.Value("email").(string)
	if email == "" {
		return nil, fmt.Errorf("email empty")
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
			ActiveAt:   time.Now().UTC().AddDate(50, 0, 0), //50 years later
			Disabled:   false,
			SubmitBy:   email,
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

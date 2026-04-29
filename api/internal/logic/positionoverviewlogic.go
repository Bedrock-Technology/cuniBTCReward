// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"
	"cuniBTCReward/model"

	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

type PositionOverviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPositionOverviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionOverviewLogic {
	return &PositionOverviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PositionOverviewLogic) PositionOverview(req *types.PositionOverviewReq) (resp []types.PositionOverviewResp, err error) {
	type aggRow struct {
		Symbol      string          `gorm:"column:symbol"`
		Amount      decimal.Decimal `gorm:"column:amount"`
		Queued      decimal.Decimal `gorm:"column:queued"`
		Earning     decimal.Decimal `gorm:"column:earning"`
		Withdrawing decimal.Decimal `gorm:"column:withdrawing"`
		Rewards     decimal.Decimal `gorm:"column:rewards"`
	}

	chainID := l.svcCtx.Config.DefaultChainId

	sql := `
	WITH latest_epoch AS (
		SELECT e1.contract, e1.lockup_start
		FROM epoches e1
		JOIN (
			SELECT contract, MAX(epoch) AS max_epoch
			FROM epoches
			WHERE chain_id = ? AND deleted_at IS NULL
			GROUP BY contract
		) e2 ON e1.contract = e2.contract AND e1.epoch = e2.max_epoch AND e1.chain_id = ? AND e1.deleted_at IS NULL
	),
	tx_agg_vault AS (
		SELECT t.contract,
			   COALESCE(SUM(t.amount), 0) AS amount,
			   COALESCE(SUM(CASE WHEN t.block_number > COALESCE(le.lockup_start,0) THEN t.amount ELSE 0 END), 0) AS queued,
			   COALESCE(SUM(CASE WHEN t.block_number <= COALESCE(le.lockup_start,0) THEN t.amount ELSE 0 END), 0) AS earning
		FROM evm_transactions t
		LEFT JOIN latest_epoch le ON le.contract = t.contract
		WHERE t.address = ? AND t.chain_id = ? AND t.deleted_at IS NULL
		  AND t.contract IN (SELECT vault FROM strategies WHERE chain_id = ? AND deleted_at IS NULL)
		GROUP BY t.contract
	),
	tx_agg_delay AS (
		SELECT t.contract,
			   COALESCE(SUM(t.amount), 0) AS amount,
			   COALESCE(SUM(CASE WHEN t.block_number > COALESCE(le.lockup_start,0) THEN t.amount ELSE 0 END), 0) AS queued,
			   COALESCE(SUM(CASE WHEN t.block_number <= COALESCE(le.lockup_start,0) THEN t.amount ELSE 0 END), 0) AS earning
		FROM evm_transactions t
		LEFT JOIN strategies s2 ON s2.delay_redeem_router = t.contract AND s2.chain_id = ? AND s2.deleted_at IS NULL
		LEFT JOIN latest_epoch le ON le.contract = s2.vault
		WHERE t.address = ? AND t.chain_id = ? AND t.deleted_at IS NULL
		GROUP BY t.contract
	),
	ad_agg AS (
		SELECT contract, COALESCE(SUM(amount),0) AS rewards
		FROM air_drop_records
		WHERE address = ? AND chain_id = ? AND deleted_at IS NULL
		GROUP BY contract
	),
	dr_agg AS (
		SELECT contract, COALESCE(SUM(amount),0) AS withdrawing
		FROM delay_redeem_records
		WHERE address = ? AND chain_id = ? AND claimed = 0 AND deleted_at IS NULL
		GROUP BY contract
	)
	SELECT s.symbol,
		   COALESCE(v.amount,0) + COALESCE(d.amount,0) + COALESCE(dr.withdrawing,0) AS amount,
		   COALESCE(v.queued,0) + COALESCE(d.queued,0) AS queued,
		   COALESCE(v.earning,0) + COALESCE(d.earning,0) AS earning,
		   COALESCE(dr.withdrawing,0) AS withdrawing,
		   COALESCE(ad.rewards,0) AS rewards
	FROM strategies s
	LEFT JOIN tx_agg_vault v ON v.contract = s.vault
	LEFT JOIN tx_agg_delay d ON d.contract = s.delay_redeem_router
	LEFT JOIN ad_agg ad ON ad.contract = s.airdrop
	LEFT JOIN dr_agg dr ON dr.contract = s.delay_redeem_router
	WHERE s.chain_id = ? AND s.deleted_at IS NULL
	`

	args := []interface{}{
		chainID, chainID, // latest_epoch
		req.Address, chainID, chainID, // tx_agg_vault (address, chain_id, subquery chain_id)
		chainID, req.Address, chainID, // tx_agg_delay (s2.chain_id, address, chain_id)
		req.Address, chainID, // ad_agg
		req.Address, chainID, // dr_agg
		chainID, // strategies
	}

	if req.Symbol != "" {
		sql += " AND s.symbol = ?"
		args = append(args, req.Symbol)
	}

	var rows []aggRow
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&rows).Error
	if err != nil {
		return
	}
	if len(rows) == 0 {
		return resp, errors.New("no strategy")
	}

	for _, r := range rows {
		resp = append(resp, types.PositionOverviewResp{
			Symbol:      r.Symbol,
			Amount:      r.Amount.Mul(decimal.New(1, -8)).String(),
			Earning:     r.Earning.Mul(decimal.New(1, -8)).String(),
			Queued:      r.Queued.Mul(decimal.New(1, -8)).String(),
			Withdrawing: r.Withdrawing.Mul(decimal.New(1, -8)).String(),
			Rewards:     r.Rewards.Mul(decimal.New(1, -8)).String(),
		})
	}
	return
}

type UserStrategyStats struct {
	Symbol      string          `json:"symbol"`
	Amount      decimal.Decimal `json:"amount"`
	Queued      decimal.Decimal `json:"queued"`
	Earning     decimal.Decimal `json:"earning"`
	Withdrawing decimal.Decimal `json:"withdrawing"`
	Rewards     decimal.Decimal `json:"rewards"`
}

func getDelayRedeemContract(symbol string, stratedy []model.Strategy) string {
	for _, v := range stratedy {
		if v.Symbol == symbol {
			return v.DelayRedeemRouter
		}
	}
	return ""
}

func (l *PositionOverviewLogic) strategyPosition(symbol string, address string, stratedy []model.Strategy) (userStatus UserStrategyStats, err error) {
	var epoch []model.Epoch
	err = l.svcCtx.Database.WithContext(l.ctx).Where("chain_id = ? AND contract = ?", l.svcCtx.Config.DefaultChainId, getVaultContract(symbol, stratedy)).
		Order("epoch desc").Limit(1).Find(&epoch).Error
	if err != nil {
		return
	}
	if len(epoch) == 0 {
		return userStatus, errors.New("no epoch")
	}
	var stats UserStrategyStats
	stats.Symbol = symbol
	//reward
	l.svcCtx.Database.WithContext(l.ctx).Model(&model.AirDropRecord{}).
		Where("address = ? AND contract = ?", address, getAirDropContract(symbol, stratedy)).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&stats.Rewards)

	//withdrawing
	l.svcCtx.Database.WithContext(l.ctx).Model(&model.DelayRedeemRecord{}).
		Where("address = ? AND contract = ? AND claimed = 0", address, getDelayRedeemContract(symbol, stratedy)).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&stats.Withdrawing)

	//queued/earning
	type TxAgg struct {
		Queued  decimal.Decimal
		Earning decimal.Decimal
		Amount  decimal.Decimal
	}
	var txAgg TxAgg
	l.svcCtx.Database.WithContext(l.ctx).Model(&model.EvmTransaction{}).
		Select(`
        SUM(CASE WHEN block_number > ? THEN amount ELSE 0 END) as queued,
        SUM(CASE WHEN block_number <= ? THEN amount ELSE 0 END) as earning,
        COALESCE(SUM(amount), 0) as amount
    `, epoch[0].LockupStart, epoch[0].LockupStart).
		Where("address = ? AND contract in ?", address,
			[]string{getVaultContract(symbol, stratedy), getDelayRedeemContract(symbol, stratedy)}).Scan(&txAgg)
	stats.Queued = txAgg.Queued
	stats.Earning = txAgg.Earning
	stats.Amount = txAgg.Amount.Add(stats.Withdrawing)

	userStatus = stats
	return
}

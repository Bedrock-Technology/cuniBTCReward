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
	// todo: add your logic here and delete this line
	var stratedy []model.Strategy
	err = l.svcCtx.Database.WithContext(l.ctx).Model(&model.Strategy{}).Where("chain_id = ?", l.svcCtx.Config.DefaultChainId).Find(&stratedy).Error
	if err != nil {
		return
	}
	if len(stratedy) == 0 {
		return resp, errors.New("no stratedy")
	}
	if req.Symbol != "" {
		s, err := getStratedy(req.Symbol, stratedy)
		if err != nil {
			return resp, errors.New("no stratedy")
		}
		stratedy = []model.Strategy{*s}
	}
	for _, v := range stratedy {
		userStatus, err := l.strategyPosition(v.Symbol, req.Address, stratedy)
		if err != nil {
			return resp, err
		}
		resp = append(resp, types.PositionOverviewResp{
			Symbol:      v.Symbol,
			Amount:      userStatus.Amount.Mul(decimal.New(1, -8)).String(),
			Earning:     userStatus.Earning.Mul(decimal.New(1, -8)).String(),
			Queued:      userStatus.Queued.Mul(decimal.New(1, -8)).String(),
			Withdrawing: userStatus.Withdrawing.Mul(decimal.New(1, -8)).String(),
			Rewards:     userStatus.Rewards.Mul(decimal.New(1, -8)).String(),
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

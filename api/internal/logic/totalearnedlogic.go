// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"
	"cuniBTCReward/model"

	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

type TotalEarnedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTotalEarnedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TotalEarnedLogic {
	return &TotalEarnedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func getAirDropContract(symbol string, stratedy []model.Strategy) string {
	for _, v := range stratedy {
		if v.Symbol == symbol {
			return v.Airdrop
		}
	}
	return ""
}

func getAirDropSymbol(airdropContract string, stratedy []model.Strategy) string {
	for _, v := range stratedy {
		if v.Airdrop == airdropContract {
			return v.Symbol
		}
	}
	return ""
}

func (l *TotalEarnedLogic) TotalEarned(req *types.TotalEarnedReq) (resp []types.TotalEarnedResp, err error) {
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

	type ContractSummary struct {
		Contract    string
		TotalAmount decimal.Decimal
	}

	var summaries []ContractSummary
	l.svcCtx.Database.WithContext(l.ctx).Model(&model.AirDropRecord{}).
		Select("contract, SUM(amount) as total_amount").
		Where("address = ?", req.Address).
		Where("contract in ?", lo.Map(stratedy, func(item model.Strategy, _ int) string {
			return item.Airdrop
		})).
		Group("contract").
		Scan(&summaries)
	totalEarned := lo.Map(summaries, func(item ContractSummary, index int) types.TotalEarnedResp {
		return types.TotalEarnedResp{
			Symbol:      getAirDropSymbol(item.Contract, stratedy),
			TotalEarned: item.TotalAmount.Mul(decimal.New(1, -8)).String(),
		}
	})

	resp = totalEarned
	return
}

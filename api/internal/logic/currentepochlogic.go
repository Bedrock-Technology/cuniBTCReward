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
	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentEpochLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentEpochLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentEpochLogic {
	return &CurrentEpochLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func getVaultContract(symbol string, stratedy []model.Strategy) string {
	for _, v := range stratedy {
		if v.Symbol == symbol {
			return v.Vault
		}
	}
	return ""
}

func getStratedy(symbol string, stratedy []model.Strategy) (*model.Strategy, error) {
	for k, v := range stratedy {
		if v.Symbol == symbol {
			return &stratedy[k], nil
		}
	}
	return nil, errors.New("not found")
}

func (l *CurrentEpochLogic) CurrentEpoch(req *types.CurrentEpochReq) (resp []types.CurrentEpochResp, err error) {
	// todo: add your logic here and delete this line
	// Use a single SQL (WITH) to fetch strategies and their latest epoch
	type epochWithSymbol struct {
		model.Epoch
		Symbol string `gorm:"column:symbol"`
	}
	chainID := l.svcCtx.Config.DefaultChainId
	var rows []epochWithSymbol
	sql := `WITH strat AS (
				SELECT vault, symbol
				FROM strategies
				WHERE chain_id = ? AND deleted_at IS NULL
			), latest AS (
				SELECT contract, MAX(epoch) AS max_epoch
				FROM epoches
				WHERE contract IN (SELECT vault FROM strat) AND deleted_at IS NULL
				GROUP BY contract
			)
			SELECT e.*, s.symbol
			FROM epoches e
			JOIN latest l ON e.contract = l.contract AND e.epoch = l.max_epoch
			JOIN strat s ON s.vault = e.contract
			WHERE e.deleted_at IS NULL`

	args := []interface{}{
		chainID,
	}

	if req.Symbol != "" {
		sql += " AND s.symbol = ?"
		args = append(args, req.Symbol)
	}
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&rows).Error
	if err != nil {
		return
	}
	if len(rows) == 0 {
		return resp, errors.New("no stratedy")
	}

	currentEpoch := lo.Map(rows, func(item epochWithSymbol, _ int) types.CurrentEpochResp {
		return types.CurrentEpochResp{
			Epoch:                 item.Epoch.Epoch,
			OperateStartTimestamp: item.Epoch.OperateStart,
			OperatePeriod:         item.Epoch.OperatePeriod,
			LockupStartTimestamp:  item.Epoch.LockupStart,
			LockupPeriod:          item.Epoch.LockupPeriod,
			Symbol:                item.Symbol,
		}
	})

	resp = currentEpoch
	return
}

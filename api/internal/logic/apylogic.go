// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"

	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

type ApyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApyLogic {
	return &ApyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApyLogic) Apy(req *types.ApyReq) (resp []types.ApyResp, err error) {
	type aggRow struct {
		Symbol string  `gorm:"column:symbol"`
		Epoch  int64   `gorm:"column:epoch"`
		Apy    float64 `gorm:"column:apy"`
	}
	sql := `
WITH latest_epoch AS (
    SELECT s.symbol, MAX(COALESCE(ae.epoch, 0)) AS epoch FROM  strategies s LEFT JOIN air_drop_epoches ae
    ON s.airdrop = ae.contract AND s.chain_id = ae.chain_id AND s.deleted_at IS NULL
    WHERE s.chain_id = ? group by s.symbol
)
SELECT le.symbol, le.epoch, COALESCE(ae.apy, 0) AS apy FROM latest_epoch le LEFT JOIN strategies s ON s.symbol = le.symbol
    AND s.deleted_at is NULL
LEFT JOIN air_drop_epoches ae ON ae.contract = s.airdrop AND ae.epoch = le.epoch WHERE ae.deleted_at IS NULL
	`
	chainID := l.svcCtx.Config.DefaultChainId
	args := []interface{}{
		chainID,
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
	for _, v := range rows {
		apyResp := types.ApyResp{
			Symbol: v.Symbol,
			Apy:    0,
		}
		apy := decimal.NewFromFloat(v.Apy)
		if apy.IsZero() {
			if v.Symbol == "suniBTC" {
				apyResp.Apy = 0.018
			} else {
				apyResp.Apy = 0.05
			}
			resp = append(resp, apyResp)
			continue
		}
		resp = append(resp, apyResp)
	}

	return
}

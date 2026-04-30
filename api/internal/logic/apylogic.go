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
	// todo: add your logic here and delete this line
	type aggRow struct {
		Symbol        string          `gorm:"column:symbol"`
		OperatePeriod uint64          `gorm:"column:operate_period"`
		LockupPeriod  uint64          `gorm:"column:lockup_period"`
		Shares        decimal.Decimal `gorm:"column:shares"`
		Rewards       decimal.Decimal `gorm:"column:rewards"`
	}

	chainID := l.svcCtx.Config.DefaultChainId
	sql := `
	WITH epoch_amount AS (
		SELECT e1.contract, e1.operate_period, e1.lockup_period, e2.shares, e2.rewards
		FROM epoches e1
		JOIN (
			SELECT s2.vault, MAX(a.epoch) AS max_epoch, COALESCE(SUM(a.shares), 0) AS shares, COALESCE(SUM(a.amount), 0) AS rewards
			FROM air_drop_records a
			LEFT JOIN strategies s2 ON s2.airdrop = a.contract AND s2.chain_id = ? AND s2.deleted_at IS NULL
			WHERE a.chain_id = ? AND a.deleted_at IS NULL
			GROUP BY s2.vault
		) e2 ON e1.contract = e2.vault AND e1.epoch = e2.max_epoch AND e1.chain_id = ? AND e1.deleted_at IS NULL
	)
	SELECT s.symbol, v.operate_period, v.lockup_period, v.shares, v.rewards
	FROM strategies s
	LEFT JOIN epoch_amount v ON v.contract = s.vault
	WHERE s.chain_id = ? AND s.deleted_at IS NULL
	`
	args := []interface{}{
		chainID, chainID, chainID, // epoch_amount
		chainID, //strategies
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
		if v.Shares.IsZero() {
			apyResp.Apy = 0.05
			resp = append(resp, apyResp)
			continue
		}
		apr := v.Rewards.Div(v.Shares).Div(decimal.NewFromUint64(v.LockupPeriod)).
			Mul(decimal.NewFromInt32(31536000))
		if apr.IsZero() {
			apyResp.Apy = 0.05
		}
		resp = append(resp, apyResp)
	}

	return
}

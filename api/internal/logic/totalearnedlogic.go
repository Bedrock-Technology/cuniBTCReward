// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"

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

func (l *TotalEarnedLogic) TotalEarned(req *types.TotalEarnedReq) (resp []types.TotalEarnedResp, err error) {
	// todo: add your logic here and delete this line
	// Use a single CTE to select strategies and aggregate air drop amounts per symbol
	type summaryRow struct {
		Symbol      string
		TotalAmount decimal.Decimal `gorm:"column:total_amount"`
	}
	chainID := l.svcCtx.Config.DefaultChainId

	var rows []summaryRow
	sql := `WITH strat AS (
				SELECT airdrop AS contract, symbol
				FROM strategies
				WHERE chain_id = ? AND deleted_at IS NULL
			)
			SELECT s.symbol, SUM(ar.amount) AS total_amount
			FROM strat s
			JOIN air_drop_records ar ON ar.contract = s.contract AND ar.address = ? AND ar.deleted_at IS NULL`

	args := []interface{}{
		chainID,
	}

	if req.Symbol != "" {
		sql += " AND s.symbol = ?"
		args = append(args, req.Symbol)
	}
	sql += " GROUP BY s.symbol"
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&rows).Error
	if err != nil {
		return
	}
	if len(rows) == 0 {
		return resp, errors.New("no stratedy")
	}

	totalEarned := lo.Map(rows, func(item summaryRow, _ int) types.TotalEarnedResp {
		return types.TotalEarnedResp{
			Symbol:      item.Symbol,
			TotalEarned: item.TotalAmount.Mul(decimal.New(1, -8)).String(),
		}
	})

	resp = totalEarned
	return
}

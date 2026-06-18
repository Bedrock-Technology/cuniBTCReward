// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"

	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

type PendingWithdrawalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPendingWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PendingWithdrawalLogic {
	return &PendingWithdrawalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PendingWithdrawalLogic) PendingWithdrawal(req *types.PendingWithdrawalReq) (resp *types.PendingWithdrawalResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId

	var requested decimal.Decimal
	sql := `SELECT COALESCE(SUM(drr.amount + drr.fee), 0) AS requested
FROM delay_redeem_records drr
JOIN strategies s ON s.delay_redeem_router = drr.contract
    AND s.chain_id = drr.chain_id AND s.deleted_at IS NULL
WHERE drr.chain_id = ? AND drr.claimed = 0 AND drr.deleted_at IS NULL
    AND s.symbol = ?`
	args := []interface{}{chainID, req.Symbol}

	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&requested).Error
	if err != nil {
		return
	}

	resp = &types.PendingWithdrawalResp{
		Requested: requested.Mul(decimal.New(1, -8)).String(),
	}
	return
}

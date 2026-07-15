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

type WithdrawalHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawalHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawalHistoryLogic {
	return &WithdrawalHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type WithdrawalHistory struct {
	Requested decimal.Decimal `gorm:"column:requested"`
	Claimed   decimal.Decimal `gorm:"column:claimed"`
	Requests  int64           `gorm:"column:requests"`
}

func (l *WithdrawalHistoryLogic) WithdrawalHistory(req *types.WithdrawalHistoryReq) (resp *types.WithdrawalHistoryResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId

	var pw WithdrawalHistory
	sql := `SELECT 
    COALESCE(SUM(drr.amount + drr.fee), 0) AS requested,
    COUNT(*) AS requests,
    COALESCE(SUM(CASE WHEN drr.claimed = 1 THEN drr.amount + drr.fee ELSE 0 END), 0) AS claimed,
FROM delay_redeem_records drr
JOIN strategies s ON s.delay_redeem_router = drr.contract
    AND s.chain_id = drr.chain_id AND s.deleted_at IS NULL
WHERE drr.chain_id = ? 
  AND drr.deleted_at IS NULL
  AND s.symbol = ?`
	args := []interface{}{chainID, req.Symbol}

	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&pw).Error
	if err != nil {
		return
	}

	resp = &types.WithdrawalHistoryResp{
		Requested: pw.Requested.Mul(decimal.New(1, -8)).String(),
		Symbol:    req.Symbol,
		Requests:  pw.Requests,
		Claimed:   pw.Claimed.Mul(decimal.New(1, -8)).String(),
	}
	return
}

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

type RewardSummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRewardSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RewardSummaryLogic {
	return &RewardSummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type rewardSummaryRow struct {
	Total     decimal.Decimal `gorm:"column:total_amount"`
	Claimed   decimal.Decimal `gorm:"column:claimed_amount"`
	UnClaimed decimal.Decimal `gorm:"column:unclaimed_amount"`
}

func (l *RewardSummaryLogic) RewardSummary(req *types.RewardSummaryReq) (resp *types.RewardSummaryResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId
	sql := `
	SELECT
    COALESCE(SUM(amount), 0) AS total_amount,
    COALESCE(SUM(CASE WHEN claimed = 1 THEN amount ELSE 0 END), 0) AS claimed_amount,
    COALESCE(SUM(CASE WHEN claimed = 0 THEN amount ELSE 0 END), 0) AS unclaimed_amount
FROM
    air_drop_records a
JOIN strategies s ON s.airdrop = a.contract WHERE s.deleted_at IS NULL AND a.deleted_at IS NULL AND s.symbol = ? AND s.chain_id = ?`
	args := []interface{}{
		req.Symbol,
		chainID,
	}
	var row rewardSummaryRow
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&row).Error
	if err != nil {
		return
	}

	resp = &types.RewardSummaryResp{
		Total:     row.Total.Mul(decimal.New(1, -8)).String(),
		Claimed:   row.Claimed.Mul(decimal.New(1, -8)).String(),
		UnClaimed: row.UnClaimed.Mul(decimal.New(1, -8)).String(),
	}

	return
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"

	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawalRequstsListReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawalRequstsListReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawalRequstsListReqLogic {
	return &WithdrawalRequstsListReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WithdrawalRequstsListReqLogic) WithdrawalRequstsListReq(req *types.WithdrawalRequstsListReq) (resp *types.WithdrawalRequstsListResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId
	if req.Offset < 0 {
		req.Offset = 0
	}
	baseSQL := `
	WITH strat AS (
        SELECT delay_redeem_router, symbol FROM strategies
        WHERE chain_id = ? AND deleted_at IS NULL AND symbol = ?
)
SELECT COALESCE(drr.amount+drr.fee,0) AS amount, CAST(UNIX_TIMESTAMP(drr.create_block_time) AS UNSIGNED) AS create_at, 
drr.claimed AS claimed FROM delay_redeem_records drr JOIN strat ON strat.delay_redeem_router = drr.contract 
WHERE address = ? AND drr.deleted_at IS NULL ORDER BY create_at DESC
`
	args := []interface{}{
		chainID, req.Symbol, // strat CTE
		req.Address,
	}
	var total int64
	countSQL := fmt.Sprintf("SELECT COUNT(*) FROM (%s)", baseSQL)
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(countSQL, args...).Scan(&total).Error
	if err != nil {
		return nil, fmt.Errorf("count failed: %v", err)
	}

	rows := []withdrawalRow{}
	dataSQL := fmt.Sprintf("%s LIMIT ? OFFSET ?", baseSQL)

	dataArgs := append(args, req.Limit, req.Offset)

	err = l.svcCtx.Database.WithContext(l.ctx).Raw(dataSQL, dataArgs...).Scan(&rows).Error
	if err != nil {
		return nil, fmt.Errorf("scan rows failed: %v", err)
	}
	data := lo.Map(rows, func(r withdrawalRow, index int) types.WithdrawalRequestsInfo {
		return types.WithdrawalRequestsInfo{
			Amount:   r.Amount.Mul(decimal.New(1, -8)).String(),
			CreateAt: r.CreateAt,
			Claimed:  r.Claimed,
		}
	})

	resp = &types.WithdrawalRequstsListResp{
		PageData: types.PageData{
			Total:  total,
			Limit:  req.Limit,
			Offset: req.Offset,
		},
		Data: data,
	}
	return
}

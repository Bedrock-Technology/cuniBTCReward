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

type UserActionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserActionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserActionsLogic {
	return &UserActionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type userActionRow struct {
	Epoch         uint64          `gorm:"column:epoch"`
	OperateStart  uint64          `gorm:"column:operate_start"`
	OperatePeriod uint64          `gorm:"column:operate_period"`
	LockupStart   uint64          `gorm:"column:lockup_start"`
	LockupPeriod  uint64          `gorm:"column:lockup_period"`
	Symbol        string          `gorm:"column:symbol"`
	Address       string          `gorm:"column:address"`
	Deposited     decimal.Decimal `gorm:"column:deposited"`
	Rewards       decimal.Decimal `gorm:"column:rewards"`
	Queued        decimal.Decimal `gorm:"column:queued"`
	ClaimAt       int64           `gorm:"column:claim_at"`
	Claimed       bool            `gorm:"column:claimed"`
}

func (l *UserActionsLogic) UserActions(req *types.UserActionsListReq) (resp *types.UserActionListResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId
	if req.Offset < 0 {
		req.Offset = 0
	}

	// Count total
	var total int64
	countSQL := `SELECT COUNT(*)
FROM air_drop_records ar
JOIN strategies s ON s.airdrop = ar.contract AND s.chain_id = ar.chain_id AND s.deleted_at IS NULL
WHERE ar.chain_id = ? AND ar.deleted_at IS NULL AND s.symbol = ? AND ar.epoch = ?`
	countArgs := []interface{}{chainID, req.Symbol, req.Epoch}
	if err = l.svcCtx.Database.WithContext(l.ctx).Raw(countSQL, countArgs...).Scan(&total).Error; err != nil {
		return
	}

	var rows []userActionRow
	sql := `SELECT ar.epoch, 0 AS operate_start, 0 AS operate_period,
       0 AS lockup_start, 0 AS lockup_period,
       s.symbol, ar.address,
       ar.shares AS deposited,
       ar.amount AS rewards,
       ar.queued AS queued,
       ar.claimed,
       COALESCE(CAST(UNIX_TIMESTAMP(ar.claim_at) AS UNSIGNED), 0) AS claim_at
FROM air_drop_records ar
JOIN strategies s ON s.airdrop = ar.contract AND s.chain_id = ar.chain_id AND s.deleted_at IS NULL
WHERE ar.chain_id = ? AND ar.deleted_at IS NULL AND s.symbol = ? AND ar.epoch = ?
ORDER BY ar.amount DESC
LIMIT ? OFFSET ?`

	args := []interface{}{
		chainID, req.Symbol, req.Epoch,
		req.Limit, req.Offset,
	}

	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&rows).Error
	if err != nil {
		return
	}

	data := make([]types.UserAction, 0, len(rows))
	for _, r := range rows {
		data = append(data, types.UserAction{
			CurrentEpochResp: types.CurrentEpochResp{
				Epoch:                 r.Epoch,
				OperateStartTimestamp: r.OperateStart,
				OperatePeriod:         r.OperatePeriod,
				LockupStartTimestamp:  r.LockupStart,
				LockupPeriod:          r.LockupPeriod,
				Symbol:                r.Symbol,
			},
			Address:   r.Address,
			Deposited: r.Deposited.Mul(decimal.New(1, -8)).String(),
			Rewards:   r.Rewards.Mul(decimal.New(1, -8)).String(),
			Queued:    r.Queued.Mul(decimal.New(1, -8)).String(),
			ClaimAt:   uint64(r.ClaimAt),
			Claimed:   r.Claimed,
		})
	}

	resp = &types.UserActionListResp{
		PageData: types.PageData{
			Total:  total,
			Limit:  req.Limit,
			Offset: req.Offset,
		},
		Data: data,
	}
	return
}

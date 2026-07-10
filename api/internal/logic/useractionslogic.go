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
JOIN epoches e ON e.contract = s.vault AND e.chain_id = ar.chain_id AND e.epoch = ar.epoch AND e.deleted_at IS NULL
WHERE ar.chain_id = ? AND ar.deleted_at IS NULL AND s.symbol = ? AND ar.epoch = ?`
	countArgs := []interface{}{chainID, req.Symbol, req.Epoch}
	if err = l.svcCtx.Database.WithContext(l.ctx).Raw(countSQL, countArgs...).Scan(&total).Error; err != nil {
		return
	}

	if total != 0 {
		var rows []userActionRow
		sql := `SELECT ar.epoch, e.operate_start, e.operate_period,
       e.lockup_start, e.lockup_period,
       s.symbol, ar.address,
       ar.shares AS deposited,
       ar.amount AS rewards,
       ar.queued AS queued,
       ar.claimed,
	   COALESCE(CAST(UNIX_TIMESTAMP(ar.claim_at) AS UNSIGNED), 0) AS claim_at
FROM air_drop_records ar
JOIN strategies s ON s.airdrop = ar.contract AND s.chain_id = ar.chain_id AND s.deleted_at IS NULL
JOIN epoches e ON e.contract = s.vault AND e.chain_id = ar.chain_id AND e.epoch = ar.epoch AND e.deleted_at IS NULL
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
	} else { //not airdrop
		var total int64
		Countsql := `
With action AS (SELECT t.address,
         COALESCE(SUM(CASE WHEN t.block_timestamp < e.lockup_start THEN t.amount ELSE 0 END), 0) AS deposited,
         COALESCE(SUM(CASE WHEN t.block_timestamp >= e.lockup_start AND t.block_timestamp < e.lockup_start + e.lockup_period THEN t.amount ELSE 0 END), 0) AS queued
FROM evm_transactions t
LEFT JOIN strategies s ON t.contract = s.vault AND s.chain_id = ? AND s.deleted_at IS NULL AND s.symbol = ?
LEFT JOIN epoches e ON e.contract = s.vault AND e.chain_id = ? AND e.deleted_at IS NULL AND epoch = ?
WHERE t.deleted_at IS NULL AND t.amount > 0
GROUP BY t.address
HAVING deposited > 0 OR queued > 0
)
SELECT COUNT(*) from action
`
		args := []interface{}{
			chainID, req.Symbol, chainID, req.Epoch,
		}
		if err = l.svcCtx.Database.WithContext(l.ctx).Raw(Countsql, args...).Scan(&total).Error; err != nil {
			return
		}
		var rows []userActionRow
		sql := `
	   SELECT t.address,
	   e.operate_start, e.operate_period,
       e.lockup_start, e.lockup_period,
       COALESCE(SUM(CASE WHEN t.block_timestamp < e.lockup_start THEN t.amount ELSE 0 END), 0) AS deposited,
       COALESCE(SUM(CASE WHEN t.block_timestamp >= e.lockup_start AND t.block_timestamp < e.lockup_start + e.lockup_period THEN t.amount ELSE 0 END), 0) AS queued
FROM evm_transactions t
LEFT JOIN strategies s ON t.contract = s.vault AND s.chain_id = ? AND s.deleted_at IS NULL AND s.symbol = ?
LEFT JOIN epoches e ON e.contract = s.vault AND e.chain_id = ? AND e.deleted_at IS NULL AND epoch = ?
WHERE t.deleted_at IS NULL AND t.amount > 0
GROUP BY t.address
HAVING deposited > 0 OR queued > 0
ORDER BY deposited DESC
LIMIT ? OFFSET ?
`
		args = append(args, req.Limit, req.Offset)
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
				ClaimAt:   uint64(0),
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
	}
	return
}

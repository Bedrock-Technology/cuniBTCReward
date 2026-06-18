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

type QueueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueueLogic {
	return &QueueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type queueRow struct {
	Amount   decimal.Decimal `gorm:"column:amount"`
	Deposits int64           `gorm:"column:deposits"`
}

func (l *QueueLogic) Queue(req *types.QueuedReq) (resp *types.QueuedResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId

	sql := `WITH latest_epoch AS (
    SELECT e.lockup_start, e.contract
    FROM epoches e
    JOIN strategies s ON s.vault = e.contract AND s.chain_id = e.chain_id AND s.deleted_at IS NULL
    WHERE e.chain_id = ? AND e.deleted_at IS NULL AND s.symbol = ?
    ORDER BY e.epoch DESC
    LIMIT 1
)
SELECT COALESCE(SUM(t.amount), 0) AS amount,
       COUNT(t.address) AS deposits
FROM latest_epoch le
LEFT JOIN evm_transactions t
    ON t.chain_id = ? AND t.contract = le.contract
    AND t.deleted_at IS NULL AND t.amount > 0
    AND t.block_timestamp >= le.lockup_start`

	args := []interface{}{chainID, req.Symbol, chainID}

	var row queueRow
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&row).Error
	if err != nil {
		return
	}

	resp = &types.QueuedResp{
		Amount:   row.Amount.Mul(decimal.New(1, -8)).String(),
		Symbol:   req.Symbol,
		Deposits: row.Deposits,
	}
	return
}

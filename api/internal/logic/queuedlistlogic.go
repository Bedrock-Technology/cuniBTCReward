// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"

	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueuedListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueuedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueuedListLogic {
	return &QueuedListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type QueuedInfoRow struct {
	Address  string          `gorm:"column:address"`
	Queued   decimal.Decimal `gorm:"column:queued"`
	CreateAt int64           `gorm:"column:block_timestamp"`
}

func (l *QueuedListLogic) QueuedList(req *types.QueuedListReq) (resp *types.QueuedListResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId
	if req.Offset < 0 {
		req.Offset = 0
	}

	var total int64
	Countsql := `
WITH latest_epoch AS (
    SELECT e.lockup_start, e.contract, e.lockup_period
    FROM epoches e
    JOIN strategies s ON s.vault = e.contract AND s.chain_id = e.chain_id AND s.deleted_at IS NULL
    WHERE e.chain_id = ? AND e.deleted_at IS NULL AND s.symbol = ?
    ORDER BY e.epoch DESC
    LIMIT 1
)
SELECT COUNT(*) AS total_count
FROM (
    SELECT t.address
    FROM evm_transactions t
    LEFT JOIN latest_epoch e ON e.contract = t.contract
    WHERE t.deleted_at IS NULL AND t.block_timestamp >= e.lockup_start AND t.block_timestamp < e.lockup_start + e.lockup_period
) AS aggregated_results;
`
	args := []interface{}{
		chainID, req.Symbol,
	}
	if err = l.svcCtx.Database.WithContext(l.ctx).Raw(Countsql, args...).Scan(&total).Error; err != nil {
		return
	}
	var rows []QueuedInfoRow
	sql := `
WITH latest_epoch AS (
    SELECT e.lockup_start, e.contract, e.lockup_period
    FROM epoches e
    JOIN strategies s ON s.vault = e.contract AND s.chain_id = e.chain_id AND s.deleted_at IS NULL
    WHERE e.chain_id = ? AND e.deleted_at IS NULL AND s.symbol = ?
    ORDER BY e.epoch DESC
    LIMIT 1
)
SELECT t.address, t.block_timestamp, t.amount
FROM evm_transactions t
LEFT JOIN latest_epoch e ON e.contract = t.contract
WHERE t.deleted_at IS NULL AND t.block_timestamp >= e.lockup_start AND t.block_timestamp < e.lockup_start + e.lockup_period
ORDER BY t.block_timestamp DESC
LIMIT ? OFFSET ?
`
	args = append(args, req.Limit, req.Offset)
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&rows).Error
	if err != nil {
		return
	}
	data := lo.Map(rows, func(x QueuedInfoRow, index int) types.QueuedInfo {
		return types.QueuedInfo{
			Address:  x.Address,
			Queued:   x.Queued.Mul(decimal.New(1, -8)).String(),
			CreateAt: x.CreateAt,
		}
	})

	resp = &types.QueuedListResp{
		PageData: types.PageData{
			Total:  total,
			Limit:  req.Limit,
			Offset: req.Offset,
		},
		Data: data,
	}
	return
}

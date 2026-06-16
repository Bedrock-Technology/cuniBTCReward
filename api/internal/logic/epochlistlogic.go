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

type EpochListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEpochListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EpochListLogic {
	return &EpochListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type epochRow struct {
	Epoch         uint64          `gorm:"column:epoch"`
	OperateStart  uint64          `gorm:"column:operate_start"`
	OperatePeriod uint64          `gorm:"column:operate_period"`
	LockupStart   uint64          `gorm:"column:lockup_start"`
	LockupPeriod  uint64          `gorm:"column:lockup_period"`
	Symbol        string          `gorm:"column:symbol"`
	Participants  int64           `gorm:"column:participants"`
	Tvl           decimal.Decimal `gorm:"column:tvl"`
	Rewards       decimal.Decimal `gorm:"column:rewards"`
	Claimed       decimal.Decimal `gorm:"column:claimed"`
}

func (l *EpochListLogic) EpochList(req *types.EpochListReq) (resp *types.EpochListResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId
	if req.Offset < 0 {
		req.Offset = 0
	}
	// Count total
	var total int64
	countSQL := `SELECT COUNT(*)
FROM epoches e
JOIN strategies s ON s.vault = e.contract AND s.chain_id = e.chain_id AND s.deleted_at IS NULL
WHERE e.chain_id = ? AND e.deleted_at IS NULL AND s.symbol = ?`
	countArgs := []interface{}{chainID, req.Symbol}
	if err = l.svcCtx.Database.WithContext(l.ctx).Raw(countSQL, countArgs...).Scan(&total).Error; err != nil {
		return
	}
	// Rebuild sql and args inline to avoid template complexity
	var rows []epochRow
	sql := `WITH strat AS (
    SELECT vault, airdrop, symbol FROM strategies
    WHERE chain_id = ? AND deleted_at IS NULL AND symbol = ?
),
epoch_tx_agg AS (
    SELECT e.contract, e.epoch,
           COUNT(DISTINCT t.address) AS participants,
           COALESCE(SUM(t.amount), 0) AS tvl
    FROM epoches e
    JOIN strat s ON s.vault = e.contract
    LEFT JOIN evm_transactions t
        ON t.chain_id = e.chain_id AND t.contract = e.contract
        AND t.deleted_at IS NULL
		-- AND t.amount > 0
        -- AND t.block_timestamp >= e.operate_start
        AND t.block_timestamp <= e.lockup_start
    WHERE e.chain_id = ? AND e.deleted_at IS NULL
    GROUP BY e.contract, e.epoch
),
airdrop_agg AS (
    SELECT s.vault,
           a.epoch,
           COALESCE(SUM(a.amount), 0) AS rewards,
           COALESCE(SUM(CASE WHEN a.claimed = 1 THEN a.amount ELSE 0 END), 0) AS claimed
    FROM air_drop_records a
    JOIN strat s ON s.airdrop = a.contract
    WHERE a.chain_id = ? AND a.deleted_at IS NULL
    GROUP BY s.vault, a.epoch
)
SELECT e.epoch, e.operate_start, e.operate_period,
       e.lockup_start, e.lockup_period, s.symbol,
       COALESCE(eta.participants, 0) AS participants,
       COALESCE(eta.tvl, 0) AS tvl,
       COALESCE(aa.rewards, 0) AS rewards,
       COALESCE(aa.claimed, 0) AS claimed
FROM epoches e
JOIN strat s ON s.vault = e.contract
LEFT JOIN epoch_tx_agg eta ON eta.contract = e.contract AND eta.epoch = e.epoch
LEFT JOIN airdrop_agg aa ON aa.vault = e.contract AND aa.epoch = e.epoch
WHERE e.chain_id = ? AND e.deleted_at IS NULL
-- ORDER BY e.epoch DESC
LIMIT ? OFFSET ?`
	args := []interface{}{
		chainID, req.Symbol, // strat CTE
		chainID, // epoch_tx_agg
		chainID, // airdrop_agg
		chainID, // main query
		req.Limit, req.Offset,
	}

	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&rows).Error
	if err != nil {
		return
	}

	data := make([]types.EpochInfo, 0, len(rows))
	for _, r := range rows {
		data = append(data, types.EpochInfo{
			CurrentEpochResp: types.CurrentEpochResp{
				Epoch:                 r.Epoch,
				OperateStartTimestamp: r.OperateStart,
				OperatePeriod:         r.OperatePeriod,
				LockupStartTimestamp:  r.LockupStart,
				LockupPeriod:          r.LockupPeriod,
				Symbol:                r.Symbol,
			},
			Participants: r.Participants,
			Tvl:          r.Tvl.Mul(decimal.New(1, -8)).String(),
			Rewards:      r.Rewards.Mul(decimal.New(1, -8)).String(),
			Claimed:      r.Claimed.Mul(decimal.New(1, -8)).String(),
		})
	}

	resp = &types.EpochListResp{
		PageData: types.PageData{
			Total:  total,
			Limit:  req.Limit,
			Offset: req.Offset,
		},
		Data: data,
	}
	return
}

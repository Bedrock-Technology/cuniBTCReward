// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActiveUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActiveUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActiveUsersLogic {
	return &ActiveUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type activeUsersRow struct {
	Participants int64 `gorm:"column:participants"`
	SignedUsers  int64 `gorm:"column:signed_users"`
}

func (l *ActiveUsersLogic) ActiveUsers(req *types.ActiveUsersReq) (resp *types.ActiveUsersResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId

	// Find symbolStatementMd5 from config
	var symbolStatementMd5 string
	for _, v := range l.svcCtx.Config.Terms {
		if v.Symbol == req.Symbol {
			symbolStatementMd5 = v.TermMd5
		}
	}

	sql := `WITH strat AS (
    SELECT vault, delay_redeem_router FROM strategies
    WHERE chain_id = ? AND deleted_at IS NULL AND symbol = ?
),
current_epoch AS (
    SELECT e.contract, e.epoch
    FROM epoches e
    JOIN strat s ON s.vault = e.contract
    WHERE e.chain_id = ? AND e.deleted_at IS NULL
    ORDER BY e.epoch DESC
    LIMIT 1
),
epoch_participants AS (
    SELECT COUNT(DISTINCT CASE WHEN total_amount != 0 THEN address END) AS participants
    FROM (
        SELECT t.address, COALESCE(SUM(t.amount), 0) AS total_amount
        FROM current_epoch ce
        JOIN epoches e ON e.contract = ce.contract AND e.epoch = ce.epoch
        JOIN strat s ON s.vault = e.contract
        LEFT JOIN evm_transactions t
            ON t.chain_id = e.chain_id
            AND t.deleted_at IS NULL
            AND (t.contract = s.vault OR t.contract = s.delay_redeem_router)
            AND t.block_timestamp <= e.lockup_start
        WHERE e.chain_id = ? AND e.deleted_at IS NULL
        GROUP BY t.address
    ) addr_agg
)
SELECT
    COALESCE((SELECT participants FROM epoch_participants), 0) AS participants,
    (SELECT COUNT(*) FROM sign_terms
     WHERE symbol = ? AND term_hash = ? AND deleted_at IS NULL) AS signed_users`

	args := []interface{}{
		chainID, req.Symbol, // strat CTE
		chainID,                        // current_epoch CTE
		chainID,                        // epoch_participants subquery
		req.Symbol, symbolStatementMd5, // sign_terms count
	}

	var row activeUsersRow
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&row).Error
	if err != nil {
		return
	}

	resp = &types.ActiveUsersResp{
		SignedUsers:  row.SignedUsers,
		Symbol:       req.Symbol,
		Participants: row.Participants,
	}
	return
}

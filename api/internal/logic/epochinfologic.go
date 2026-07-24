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

type EpochInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEpochInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EpochInfoLogic {
	return &EpochInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EpochInfoLogic) EpochInfo(req *types.EpochInfoReq) (resp *types.EpochInfo, err error) {
	chainID := l.svcCtx.Config.DefaultChainId
	var rows []epochRow
	sql := `WITH strat AS (
    SELECT vault, airdrop, delay_redeem_router, symbol FROM strategies
    WHERE chain_id = ? AND deleted_at IS NULL AND symbol = ?
),
top_epoches AS (
    SELECT e.*
    FROM epoches e
    JOIN strat s ON s.vault = e.contract
    WHERE e.chain_id = ?
      AND e.deleted_at IS NULL AND e.epoch = ?
),
epoch_addr_sum AS (
    SELECT e.contract, e.epoch, t.address,
           COALESCE(SUM(t.amount), 0) AS total_amount
    FROM top_epoches e
    JOIN strat s ON s.vault = e.contract
    LEFT JOIN evm_transactions t
        ON t.chain_id = e.chain_id
        AND t.deleted_at IS NULL
        AND t.contract IN (s.vault, s.delay_redeem_router)
        AND t.block_timestamp < e.lockup_start + e.lockup_period
    WHERE e.chain_id = ? AND e.deleted_at IS NULL
    GROUP BY e.contract, e.epoch, t.address
),
epoch_tx_agg AS (
    SELECT contract, epoch,
           COUNT(DISTINCT CASE WHEN total_amount != 0 THEN address END) AS participants,
           COALESCE(SUM(total_amount), 0) AS tvl_trans
    FROM epoch_addr_sum
    GROUP BY contract, epoch
),
epoch_unclaimed AS (
    SELECT te.contract,
		   te.epoch,
           COALESCE(SUM(drr.amount + drr.fee), 0) AS unclaimed_redeem
    FROM delay_redeem_records drr
    JOIN strat s ON s.delay_redeem_router = drr.contract
	JOIN top_epoches te ON te.contract = s.vault
    WHERE drr.chain_id = ? 
      AND drr.deleted_at IS NULL
	  AND drr.create_block_time < FROM_UNIXTIME(te.lockup_start + te.lockup_period)
	  AND (
	      drr.claimed = 0
	      OR drr.claim_at > FROM_UNIXTIME(te.lockup_start + te.lockup_period)
		  )
    GROUP BY te.contract, te.epoch
),
airdrop_agg AS (
    SELECT s.vault,
           a.epoch,
           COALESCE(SUM(a.amount), 0) AS rewards,
           COALESCE(SUM(CASE WHEN a.claimed = 1 THEN a.amount ELSE 0 END), 0) AS claimed
    FROM air_drop_records a
    JOIN strat s ON s.airdrop = a.contract
	JOIN top_epoches te ON te.contract = s.vault AND te.epoch = a.epoch
    WHERE a.chain_id = ? AND a.deleted_at IS NULL
    GROUP BY s.vault, a.epoch
)
SELECT e.epoch, e.operate_start, e.operate_period,
       e.lockup_start, e.lockup_period, s.symbol,
       COALESCE(eta.participants, 0) AS participants,
       COALESCE(eta.tvl_trans, 0) + COALESCE(u.unclaimed_redeem, 0) AS tvl,
       COALESCE(aa.rewards, 0) AS rewards,
       COALESCE(aa.claimed, 0) AS claimed,
       COALESCE(ae.apy, 0) AS apy,
       ae.root,
       ae.merkle_root,
	   ae.token,
	   ae.created_at,
	   ae.submit_by,
	   ae.creator_at,
	   ae.creator
FROM top_epoches e
JOIN strat s ON s.vault = e.contract
LEFT JOIN epoch_tx_agg eta ON eta.contract = e.contract AND eta.epoch = e.epoch
LEFT JOIN epoch_unclaimed u ON u.contract = e.contract AND u.epoch = e.epoch
LEFT JOIN airdrop_agg aa ON aa.vault = e.contract AND aa.epoch = e.epoch
LEFT JOIN air_drop_epoches ae ON ae.contract = s.airdrop AND ae.epoch = e.epoch AND ae.chain_id = e.chain_id AND ae.deleted_at IS NULL
WHERE e.chain_id = ? AND e.deleted_at IS NULL
ORDER BY e.epoch DESC
`
	args := []interface{}{
		chainID, req.Symbol, // strat CTE
		chainID,
		req.Epoch,
		chainID, // epoch_addr_sum
		chainID, // epoch_unclaimed
		chainID, // airdrop_agg
		chainID, // main query
	}

	err = l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...).Scan(&rows).Error
	if err != nil {
		return
	}
	for _, r := range rows {
		resp = &types.EpochInfo{
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
			Apy:          r.Apy,
			Root:         r.Root,
			MerkleRoot:   r.MerkleRoot,
			RewardToken:  r.RewardToken,
			SubmitAt: func() int64 {
				if r.SubmitAt.Valid {
					return r.SubmitAt.Time.Unix()
				} else {
					return 0
				}
			}(),
			SubmitBy: r.SubmitBy,
		}
	}

	return
}

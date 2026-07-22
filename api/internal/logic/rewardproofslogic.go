// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"encoding/json"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"

	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

type RewardProofsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRewardProofsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RewardProofsLogic {
	return &RewardProofsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type rewardProofRow struct {
	Epoch   uint64          `gorm:"column:epoch"`
	Claimed bool            `gorm:"column:claimed"`
	Proof   string          `gorm:"column:proof"`
	Amount  decimal.Decimal `gorm:"column:amount"`
}

func (l *RewardProofsLogic) RewardProofs(req *types.RewardProofsReq) (resp *types.RewardProofsResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId
	sql := `WITH strat AS (
    SELECT airdrop, symbol FROM strategies
    WHERE chain_id = ? AND deleted_at IS NULL AND symbol = ?
),
adr AS (
SELECT claimed, epoch, proof, contract, amount FROM air_drop_records WHERE chain_id = ? AND address = ? AND deleted_at IS NULL
),
ade AS (
SELECT active_at, contract, epoch FROM air_drop_epoches WHERE chain_id = ? AND deleted_at IS NULL
)`
	args := []interface{}{
		chainID,
		req.Symbol,
		chainID,
		req.Address,
		chainID,
	}
	var totalAmount decimal.Decimal
	totalSql := sql + " SELECT COALESCE(SUM(amount), 0) AS totalAmount FROM adr JOIN strat ON strat.airdrop  = adr.contract JOIN ade ON ade.contract = strat.airdrop AND adr.epoch = ade.epoch WHERE active_at < NOW() ORDER BY adr.epoch DESC"
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(totalSql, args...).Scan(&totalAmount).Error
	if err != nil {
		return
	}
	var rows []rewardProofRow
	proofSql := sql + " SELECT claimed, adr.epoch AS epoch, proof, amount FROM adr JOIN strat ON strat.airdrop  = adr.contract JOIN ade ON ade.contract = strat.airdrop AND adr.epoch = ade.epoch WHERE active_at < NOW() AND amount > 0 ORDER BY adr.epoch DESC"
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(proofSql, args...).Scan(&rows).Error
	if err != nil {
		return
	}
	resp = &types.RewardProofsResp{
		TotalAmount: totalAmount.Mul(decimal.New(1, -8)).String(),
	}
	proofs := lo.Map(rows, func(row rewardProofRow, index int) types.RewardProofs {
		proof := []string{}
		_ = json.Unmarshal([]byte(row.Proof), &proof)
		rewardProofs := types.RewardProofs{
			Epoch:   row.Epoch,
			Amount:  row.Amount.Mul(decimal.New(1, -8)).String(),
			Claimed: row.Claimed,
			Proofs:  proof,
		}
		return rewardProofs
	})

	resp.RewardProofs = proofs
	return
}

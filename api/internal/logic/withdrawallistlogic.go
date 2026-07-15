// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"

	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawalListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawalListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawalListLogic {
	return &WithdrawalListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type withdrawalRow struct {
	Address  string          `gorm:"column:address"`
	Amount   decimal.Decimal `gorm:"column:amount"`
	Epoch    uint64          `gorm:"column:epoch"`
	CreateAt int64           `gorm:"column:create_at"`
	ClaimAt  int64           `gorm:"column:claim_at"`
	Claimed  bool            `gorm:"column:claimed"`
}

func (l *WithdrawalListLogic) WithdrawalList(req *types.WithdrawalListReq) (resp *types.WithdrawalListResp, err error) {
	chainID := l.svcCtx.Config.DefaultChainId
	if req.Offset < 0 {
		req.Offset = 0
	}
	// Count total
	var total int64
	sql := `
	WITH strat AS (
    SELECT vault, airdrop, delay_redeem_router, symbol FROM strategies
    WHERE chain_id = ? AND deleted_at IS NULL AND symbol = ?
),
epoches AS (
    SELECT e.*
    FROM epoches e
    JOIN strat s ON s.vault = e.contract
    WHERE e.chain_id = ?
      AND e.deleted_at IS NULL
)
SELECT drr.address, COALESCE(drr.amount+drr.fee,0) AS amount, e.epoch, drr.create_block_time AS create_at, drr.claim_at, drr.claimed FROM delay_redeem_records drr INNER JOIN epoches e
    ON UNIX_TIMESTAMP(drr.create_block_time) >= e.operate_start AND UNIX_TIMESTAMP(drr.create_block_time) < e.lockup_start + e.lockup_period
	WHERE drr.deleted_at IS NULL
	`
	args := []interface{}{
		chainID, req.Symbol, // strat CTE
		chainID,
	}
	tx := l.svcCtx.Database.WithContext(l.ctx).Raw(sql, args...)
	if req.Address != "" {
		tx.Where("drr.address = ?", req.Address)
	}
	if req.Start != 0 {
		tx.Where("UNIX_TIMESTAMP(drr.create_block_time) >= ?", req.Start)
	}
	if req.End != 0 {
		tx.Where("UNIX_TIMESTAMP(drr.create_block_time) < ?", req.End)
	}
	if req.Epoch != "" {
		epoch, err := strconv.Atoi(req.Epoch)
		if err != nil {
			return nil, fmt.Errorf("epoch not number")
		}
		tx.Where("e.epoch = ?", epoch)
	}

	switch req.Status {
	case "claimed":
		tx.Where("drr.claimed = ?", true)
	case "unClaimed":
		tx.Where("drr.claimed = ?", false)
	case "coolingDown":
		tx.Where("drr.claimed = ?", false).Where("UNIX_TIMESTAMP(drr.create_block_time) < ?", time.Now().UTC().AddDate(0, 0, -7).Unix())
	default:
	}
	tx.Limit(req.Limit).Offset(req.Offset)

	rows := []withdrawalRow{}

	err = tx.Count(&total).Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	data := make([]types.WithdrawalInfo, 0, len(rows))
	for _, r := range rows {
		data = append(data, types.WithdrawalInfo{
			Address:  r.Address,
			Amount:   r.Amount.Mul(decimal.New(1, -8)).String(),
			Epoch:    r.Epoch,
			CreateAt: r.CreateAt,
			ClaimAt:  r.ClaimAt,
			Claimed:  r.Claimed,
		})
	}
	resp = &types.WithdrawalListResp{
		PageData: types.PageData{
			Total:  total,
			Limit:  req.Limit,
			Offset: req.Offset,
		},
		Data: data,
	}
	return
}

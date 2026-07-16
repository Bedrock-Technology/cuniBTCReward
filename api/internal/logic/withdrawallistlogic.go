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

	baseSQL := `
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
    SELECT drr.address AS address, COALESCE(drr.amount+drr.fee,0) AS amount, e.epoch AS epoch, CAST(UNIX_TIMESTAMP(drr.create_block_time) AS UNSIGNED) AS create_at,AST(UNIX_TIMESTAMP(drr.claim_at) AS UNSIGNED) AS claim_at, drr.claimed AS claimed 
    FROM delay_redeem_records drr 
    INNER JOIN epoches e ON UNIX_TIMESTAMP(drr.create_block_time) >= e.operate_start 
                        AND UNIX_TIMESTAMP(drr.create_block_time) < e.lockup_start + e.lockup_period
    WHERE drr.deleted_at IS NULL`

	args := []interface{}{
		chainID, req.Symbol, // strat CTE
		chainID, // epoches CTE
	}

	dynamicWhere := ""
	var dynamicArgs []interface{}

	if req.Address != "" {
		dynamicWhere += " AND drr.address = ?"
		dynamicArgs = append(dynamicArgs, req.Address)
	}
	if req.Start != 0 {
		dynamicWhere += " AND UNIX_TIMESTAMP(drr.create_block_time) >= ?"
		dynamicArgs = append(dynamicArgs, req.Start)
	}
	if req.End != 0 {
		dynamicWhere += " AND UNIX_TIMESTAMP(drr.create_block_time) < ?"
		dynamicArgs = append(dynamicArgs, req.End)
	}
	if req.Epoch != "" {
		epoch, err := strconv.Atoi(req.Epoch)
		if err != nil {
			return nil, fmt.Errorf("epoch not number")
		}
		dynamicWhere += " AND e.epoch = ?"
		dynamicArgs = append(dynamicArgs, epoch)
	}

	switch req.Status {
	case "claimed":
		dynamicWhere += " AND drr.claimed = ?"
		dynamicArgs = append(dynamicArgs, true)
	case "unClaimed":
		dynamicWhere += " AND drr.claimed = ?"
		dynamicArgs = append(dynamicArgs, false)
	case "coolingDown":
		dynamicWhere += " AND drr.claimed = ? AND UNIX_TIMESTAMP(drr.create_block_time) < ?"
		dynamicArgs = append(dynamicArgs, false, time.Now().UTC().AddDate(0, 0, -7).Unix())
	}

	fullArgs := append(args, dynamicArgs...)

	var total int64
	countSQL := fmt.Sprintf("SELECT COUNT(*) FROM (%s %s) AS temp_count_table", baseSQL, dynamicWhere)
	err = l.svcCtx.Database.WithContext(l.ctx).Raw(countSQL, fullArgs...).Scan(&total).Error
	if err != nil {
		return nil, fmt.Errorf("count failed: %v", err)
	}

	if total == 0 {
		return &types.WithdrawalListResp{
			PageData: types.PageData{
				Total:  0,
				Limit:  req.Limit,
				Offset: req.Offset,
			},
			Data: []types.WithdrawalInfo{},
		}, nil
	}

	rows := []withdrawalRow{}
	dataSQL := fmt.Sprintf("%s %s ORDER BY epoch DESC LIMIT ? OFFSET ?", baseSQL, dynamicWhere)

	dataArgs := append(fullArgs, req.Limit, req.Offset)

	err = l.svcCtx.Database.WithContext(l.ctx).Raw(dataSQL, dataArgs...).Scan(&rows).Error
	if err != nil {
		return nil, fmt.Errorf("scan rows failed: %v", err)
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
	return resp, nil
}

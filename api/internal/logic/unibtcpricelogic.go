// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"cuniBTCReward/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UniBTCPriceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUniBTCPriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UniBTCPriceLogic {
	return &UniBTCPriceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UniBTCPriceLogic) UniBTCPrice() (resp uint64, err error) {
	// todo: add your logic here and delete this line
	resp = l.svcCtx.UniBtcPriceCron.GetUniBTCPrice()
	return
}

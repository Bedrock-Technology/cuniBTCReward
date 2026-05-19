// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"
	"cuniBTCReward/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignTermsIteamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignTermsIteamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignTermsIteamLogic {
	return &SignTermsIteamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignTermsIteamLogic) SignTermsIteam(req *types.SignTermsIteamReq) (resp *types.SignTermsIteamResp, err error) {
	// todo: add your logic here and delete this line
	signTerms := []model.SignTerms{}
	err = l.svcCtx.Database.WithContext(l.ctx).Where("address = ?", req.Address).Order("nonce desc").Limit(1).Error
	if err != nil {
		return nil, err
	}
	if len(signTerms) == 0 {
		return nil, nil
	}
	iteam := signTerms[0]
	resp = &types.SignTermsIteamResp{
		Message: types.Message{
			Address:    iteam.Address,
			Nonce:      iteam.Nonce,
			Content:    iteam.Content,
			ExpireTime: uint64(iteam.ExpireTime.Unix()),
		},
	}
	return
}

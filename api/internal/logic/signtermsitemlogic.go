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

type SignTermsItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignTermsItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignTermsItemLogic {
	return &SignTermsItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignTermsItemLogic) SignTermsItem(req *types.SignTermsItemReq) (resp *types.SignTermsItemResp, err error) {
	// todo: add your logic here and delete this line
	signTerms := []model.SignTerms{}
	err = l.svcCtx.Database.WithContext(l.ctx).Where("address = ?", req.Address).Order("nonce desc").Limit(1).Find(&signTerms).Error
	if err != nil {
		return nil, err
	}
	if len(signTerms) == 0 {
		return nil, nil
	}
	iteam := signTerms[0]
	resp = &types.SignTermsItemResp{
		Message: types.Message{
			Address:    iteam.Address,
			Nonce:      iteam.Nonce,
			Content:    iteam.Content,
			ExpireTime: uint64(iteam.ExpireTime.Unix()),
		},
	}
	return
}

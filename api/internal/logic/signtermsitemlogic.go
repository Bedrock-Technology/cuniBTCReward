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
	var symbolStatementMd5 string
	for _, v := range l.svcCtx.Config.Terms {
		if v.Symbol == req.Symbol {
			symbolStatementMd5 = v.TermMd5
		}
	}
	signTerms := []model.SignTerms{}
	err = l.svcCtx.Database.WithContext(l.ctx).Where("address = ?", req.Address).
		Where("symbol = ?", req.Symbol).Where("term_hash = ?", symbolStatementMd5).Limit(1).Find(&signTerms).Error
	if err != nil {
		return nil, err
	}
	if len(signTerms) == 0 {
		return nil, nil
	}
	item := signTerms[0]
	resp = &types.SignTermsItemResp{
		Address:   item.Address,
		Message:   item.Message,
		Symbol:    req.Symbol,
		Signature: item.Signature,
	}
	return
}

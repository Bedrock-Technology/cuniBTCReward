// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"
	"cuniBTCReward/model"

	"github.com/spruceid/siwe-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignTermsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignTermsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignTermsLogic {
	return &SignTermsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignTermsLogic) SignTerms(req *types.SignTermsReq) (resp *types.SignTermsResp, err error) {
	// todo: add your logic here and delete this line
	message, err := siwe.ParseMessage(req.Message)
	if err != nil {
		return
	}
	if *message.GetStatement() == "" || message.GetAddress().String() == "" {
		return resp, fmt.Errorf("not contain, statement or address")
	}
	statementMd5 := md5.Sum([]byte(*message.GetStatement()))
	statementMd5Str := hex.EncodeToString(statementMd5[:])

	var symbolStatementMd5 string
	for _, v := range l.svcCtx.Config.Terms {
		if v.Symbol == req.Symbol {
			symbolStatementMd5 = v.TermMd5
		}
	}

	if statementMd5Str != symbolStatementMd5 {
		l.Infof("address: %s, symbolStatementMd5: %s, statementMd5Str: %s, message: %s",
			message.GetAddress().String(), symbolStatementMd5, statementMd5Str, req.Message)
		l.Errorf("address: %s, signTerms error", message.GetAddress().String())
		return nil, fmt.Errorf("statement not equal")
	}

	_, err = message.VerifyEIP191(req.Signature)
	if err != nil {
		return
	}
	//save into db
	term := model.SignTerms{
		Address:   message.GetAddress().String(),
		Symbol:    req.Symbol,
		TermHash:  statementMd5Str,
		Message:   req.Message,
		Signature: req.Signature,
	}
	if err := l.svcCtx.Database.WithContext(l.ctx).Save(&term).Error; err != nil {
		return nil, err
	}

	resp = &types.SignTermsResp{
		Address:  message.GetAddress().String(),
		TermsMd5: statementMd5Str,
		Symbol:   req.Symbol,
	}
	return
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"
	"cuniBTCReward/model"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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
	var message types.Message
	err = json.Unmarshal([]byte(req.Message), &message)
	if err != nil {
		return nil, fmt.Errorf("message unmarshal:%v", err)
	}
	if message.Address == "" || message.Content == "" || message.ExpireTime == 0 {
		return nil, fmt.Errorf("message error")
	}
	if !verifySig(common.HexToAddress(message.Address).String(), req.Signature, []byte(req.Message)) {
		return nil, fmt.Errorf("sign error")
	}
	//write to db
	signTerms := model.SignTerms{
		Address:    common.HexToAddress(message.Address).String(),
		Nonce:      message.Nonce,
		ExpireTime: time.Unix(int64(message.ExpireTime), 0),
		Content:    req.Message,
	}
	err = l.svcCtx.Database.WithContext(l.ctx).Create(&signTerms).Error
	if err != nil {
		return nil, err
	}

	resp = &types.SignTermsResp{
		Message: message,
	}
	return
}

func verifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	return from == recoveredAddr.Hex()
}

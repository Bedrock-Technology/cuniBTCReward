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
	recoverAddress, err := l.evmSignVerify(req.Message, req.Signature)
	if err != nil {
		return nil, err
	}
	if recoverAddress != common.HexToAddress(message.Address).String() {
		return nil, fmt.Errorf("sign error")
	}
	//write to db
	signTerms := model.SignTerms{
		Address:    recoverAddress,
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

func (l *SignTermsLogic) evmSignVerify(message, sig string) (string, error) {
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	messageHash := crypto.Keccak256Hash([]byte(prefixedMessage))
	signature, err := hexutil.Decode(sig)
	if err != nil {
		return "", err
	}
	// Adjust the recovery ID (v) if needed (e.g., if v is 27/28)
	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}
	// 3. Recover the public key
	pubKeyBytes, err := crypto.Ecrecover(messageHash.Bytes(), signature)
	if err != nil {
		return "", err
	}
	publicKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return "", err
	}
	// 4. Derive the address and verify
	recoveredAddress := crypto.PubkeyToAddress(*publicKey)
	l.Infof("recoverAddress:%v", recoveredAddress)
	return recoveredAddress.String(), nil
}

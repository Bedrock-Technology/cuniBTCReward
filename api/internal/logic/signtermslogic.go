// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"
	"cuniBTCReward/model"

	"github.com/spruceid/siwe-go"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignTermsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var (
	once    sync.Once
	limiter *limit.PeriodLimit
)

func NewSignTermsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignTermsLogic {
	once.Do(func() {
		limiter = limit.NewPeriodLimit(60, 3, svcCtx.Redis, "cuniBTC:signTerm:rate:")
	})
	return &SignTermsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func getRealIP(r *http.Request) string {
	// 1. Check the standard Cloudflare header
	if ip := r.Header.Get("CF-Connecting-IP"); ip != "" {
		return ip
	}

	// 2. Check the standard proxy chain header if Cloudflare header is missing
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// X-Forwarded-For can contain multiple IPs separated by commas.
		// The first IP is always the original client.
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}

	// 3. Fallback to direct remote address if not proxied
	remoteIP := r.RemoteAddr
	if strings.Contains(remoteIP, ":") {
		parts := strings.Split(remoteIP, ":")
		remoteIP = strings.Join(parts[:len(parts)-1], ":")
	}
	return remoteIP
}

func (l *SignTermsLogic) SignTerms(req *types.SignTermsReq, r *http.Request) (resp *types.SignTermsResp, err error) {
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
		l.Infof("address: %s, symbolStatementMd5: %s, statementMd5Str: %s, message: %s, statement: %s",
			message.GetAddress().String(), symbolStatementMd5, statementMd5Str, req.Message, *message.GetStatement())
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
	} else {
		//limiter
		realIp := getRealIP(r)
		l.Infof("signTerm ip:%s", realIp)
		code, err := limiter.TakeCtx(r.Context(), realIp)
		if err != nil {
			l.Error("signTerm limiter error")
		}

		switch code {
		case limit.OverQuota:
			return nil, fmt.Errorf("too many requests")
		}
	}

	resp = &types.SignTermsResp{
		Address:  message.GetAddress().String(),
		TermsMd5: statementMd5Str,
		Symbol:   req.Symbol,
	}
	return
}

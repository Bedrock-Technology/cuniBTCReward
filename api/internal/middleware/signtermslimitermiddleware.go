// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package middleware

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type SignTermsLimiterMiddleware struct {
	limiter *limit.PeriodLimit
}

func NewSignTermsLimiterMiddleware(rds *redis.Redis) *SignTermsLimiterMiddleware {
	lmt := limit.NewPeriodLimit(60, 10, rds, "cuniBTC:signTerm:rate:")
	return &SignTermsLimiterMiddleware{
		limiter: lmt,
	}
}

func (m *SignTermsLimiterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		realIp := getRealIP(r)
		logx.WithContext(r.Context()).Infof("signTerm ip:%s", realIp)
		code, err := m.limiter.TakeCtx(r.Context(), realIp)
		if err != nil {
			logx.WithContext(r.Context()).Error("signTerm limiter error")
			next(w, r)
			return
		}
		switch code {
		case limit.OverQuota:
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}
		next(w, r)
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

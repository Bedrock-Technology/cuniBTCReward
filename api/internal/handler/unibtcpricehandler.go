package handler

import (
	"cuniBTCReward/api/internal/logic"
	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/response"
	"net/http"
)

func UniBTCPriceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUniBTCPriceLogic(r.Context(), svcCtx)
		resp, err := l.UniBTCPrice()
		response.Response(w, resp, err)

	}
}

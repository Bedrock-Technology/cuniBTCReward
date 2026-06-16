package handler

import (
	"cuniBTCReward/api/internal/logic"
	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"
	"cuniBTCReward/api/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EpochListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EpochListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEpochListLogic(r.Context(), svcCtx)
		resp, err := l.EpochList(&req)
		response.Response(w, resp, err)

	}
}

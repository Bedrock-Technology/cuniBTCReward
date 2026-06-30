package handler

import (
	"net/http"

	"github.com/swaggest/swgui/v5emb"
	"github.com/zeromicro/go-zero/rest"
)

func RegisterSwaggerHandlers(server *rest.Server) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/swagger/",
				Handler: SwaggerHandler(),
			},
		},
	)
}

func SwaggerHandler() http.HandlerFunc {
	swagger := v5emb.New(
		"cuniBTC",
		"/docs/cuniBTCReward.json",
		"/api/swagger/",
	)
	return func(w http.ResponseWriter, r *http.Request) {
		swagger.ServeHTTP(w, r)
	}
}

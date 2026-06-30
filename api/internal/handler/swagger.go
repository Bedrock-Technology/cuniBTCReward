package handler

import (
	"net/http"

	"github.com/swaggest/swgui/v5cdn"
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
	swagger := v5cdn.New(
		"cuniBTC",
		"/docs/cuniBTCReward.json",
		"",
	)
	return func(w http.ResponseWriter, r *http.Request) {
		swagger.ServeHTTP(w, r)
	}
}

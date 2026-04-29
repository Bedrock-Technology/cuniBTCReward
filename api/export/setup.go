package setup

import (
	exportconfig "cuniBTCReward/api/export/config"
	"cuniBTCReward/api/internal/config"
	"cuniBTCReward/api/internal/handler"
	"cuniBTCReward/api/internal/svc"
	"unsafe"

	"github.com/zeromicro/go-zero/rest"
)

func Setup(server *rest.Server, c exportconfig.Config) {
	config := (*config.Config)(unsafe.Pointer(&c))
	ctx := svc.NewServiceContext(*config)
	handler.RegisterHandlers(server, ctx)
	handler.RegisterSwaggerHandlers(server)
	server.PrintRoutes()
}

package internal

import (
	"time"

	"github.com/Dmytro-Kucherenko/smartner-users-service/docs"
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/common/config"
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	validator "github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/adapters/playground"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	ShutdownTimeout time.Duration = 10 * time.Second
	DBTimeout       time.Duration = 1 * time.Second
)

func addDocs() {
	host := config.AppHost()
	path := config.AppBasePath()
	protocol := config.AppProtocol()

	docs.SwaggerInfo.Title = "Users API"
	docs.SwaggerInfo.Description = "API server to handle users requests."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = path
	docs.SwaggerInfo.Schemes = []string{protocol}
}

func Init(logger types.Logger, meta server.RequestMeta) (options adapter.StartupOptions, err error) {
	err = config.Load()
	if err != nil {
		return
	}

	connection := config.DBConnection()
	port := config.AppPort()
	clientURL := config.ClientURL()
	isProd := config.IsProd()

	db, err := server.ConnectSQL(connection, DBTimeout)
	if err != nil {
		return
	}

	err = validator.TryRegister(binding.Validator.Engine())
	if err != nil {
		return
	}

	addDocs()
	router, httpServer := adapter.CreateRouter(port, isProd, clientURL)
	api := adapter.CreateRoutes(router, "/users", logger)
	modules.Init(api, db, meta)

	return adapter.StartupOptions{
		Router: router,
		StartupOptions: server.StartupOptions{
			Server:          httpServer,
			ShutdownTimeout: ShutdownTimeout,
		},
	}, nil
}

package internal

import (
	"net/http"
	"time"

	"github.com/dmytro-kucherenko/smartner-users-service/docs"
	"github.com/dmytro-kucherenko/smartner-users-service/internal/common/config"
	"github.com/dmytro-kucherenko/smartner-users-service/internal/modules"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	schema "github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/adapters/playground"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	adapterGin "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
	adapterGRPC "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/grpc"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/grpc/interceptors"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/multiplexer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Make default or so, log grpc interceptor
// Layered Architecture: repos, services, apis
// Remove parsers, make  manual mappers inside api/repo
const (
	dbTimeout time.Duration = 1 * time.Second
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

func New() (app *modules.App, err error) {
	err = config.Load()
	if err != nil {
		return
	}

	err = schema.Init()
	if err != nil {
		return
	}

	connection := config.DBConnection()
	db, err := server.ConnectSQL(connection, dbTimeout)
	if err != nil {
		return
	}

	UserConnection := "localhost:8000" // from config
	userConn, err := grpc.NewClient(UserConnection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	return modules.NewApp(db, userConn), nil
}

func InitREST(app *modules.App) *http.Server {
	logger := log.New("init-rest")
	clientURL := config.ClientURL()
	isProd := config.IsProd()

	addDocs()
	router, server := adapterGin.CreateRouter(isProd, clientURL)
	api := adapterGin.CreateRoutes(router, "/user", logger)
	adapterGin.InitModules(api, app)

	return server
}

func InitGRPC(app *modules.App) *grpc.Server {
	config := adapterGRPC.GetConfig(app)
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptors.OptionsUnary(),
		interceptors.ConfigUnary(config),
		interceptors.ValidateUnary()),
	)

	adapterGRPC.InitModules(server, app)

	return server
}

func Init() error {
	app, err := New()
	if err != nil {
		return err
	}

	port := config.AppPort()
	service, err := multiplexer.NewService(port)
	if err != nil {
		return err
	}

	err = service.WithGRPC(InitGRPC(app)).WithHTTP(InitREST(app)).Serve()
	if err != nil {
		return err
	}

	return nil
}

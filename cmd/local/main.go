package main

import (
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/google/uuid"
)

func InitWithMeta(create func(logger types.Logger, meta server.RequestMeta) (internal.StartupOptionsInternal, error), meta server.RequestMeta) {
	logger := log.New("Init")
	options, err := create(logger, meta)
	if err != nil {
		panic(err.Error())
	}

	if options.OnlyConfig {
		logger.Info("app was configured")

		return
	}

	err = server.ServeGracefully(options.Server, logger, options.ShutdownTimeout)
	if err != nil {
		panic(err.Error())
	}
}

func Init(create func(logger types.Logger, meta server.RequestMeta) (internal.StartupOptionsInternal, error)) {
	id, _ := uuid.Parse("451f4f07-5140-456f-9ffc-4751a808f45f")
	meta := server.RequestMeta{
		Session: &server.Session{
			UserID: id,
		},
	}

	InitWithMeta(create, meta)
}

// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	JWTAuth
// @in							header
// @name						Authorization
// @description				JWT authorization guard
func main() {
	Init(func(logger types.Logger, meta server.RequestMeta) (internal.StartupOptionsInternal, error) {
		options, err := internal.Init(logger, meta)

		return options.StartupOptionsInternal, err
	})
}

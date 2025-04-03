package main

import (
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	startup "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/startups/local"
)

// @version					1.0
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	JWTAuth
// @in							Cookie
// @name						Authorization
// @description				JWT authorization guard
func main() {
	startup.Init(func(logger types.Logger, meta server.RequestMeta) (server.StartupOptions, error) {
		options, err := internal.Init(logger, meta)

		return options.StartupOptions, err
	})
}

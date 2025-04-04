package main

import (
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal"
	startup "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin/startups/lambda"
)

// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	JWTAuth
// @in							header
// @name						Authorization
// @description				JWT authorization guard
func main() {
	startup.Init(internal.Init)
}

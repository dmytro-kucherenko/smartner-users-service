package main

import (
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal"
	startup "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin/startups/lambda"
)

// @version					1.0
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	JWTAuth
// @in							Cookie
// @name						Authorization
// @description				JWT authorization guard
func main() {
	startup.Init(internal.Init)
}

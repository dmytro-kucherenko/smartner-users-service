package main

import (
	"github.com/dmytro-kucherenko/smartner-users-service/internal"
)

// @version					1.0
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	JWTAuth
// @in							Cookie
// @name						Authorization
// @description				JWT authorization guard
func main() {
	err := internal.Init()
	if err != nil {
		panic(err.Error())
	}
}

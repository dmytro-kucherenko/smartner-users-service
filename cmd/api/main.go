package main

import "github.com/Dmytro-Kucherenko/smartner-users-service/internal"

// @title						Users API
// @version					1.0
// @description				This is a users service example.
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8000
// @BasePath					/api/v1
// @securityDefinitions.apikey	JWTAuth
// @in							header
// @name						Authorization
// @description				Bearer JWT authorization guard
func main() {
	internal.Init()
}

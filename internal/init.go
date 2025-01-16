package internal

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/Dmytro-Kucherenko/smartner-users-service/docs"
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/common/config"
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/validations"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const SHUTDOWN_TIMEOUT time.Duration = 10 * time.Second

func Init() {
	logger := log.New("Init")

	connection := config.DBConnection()
	port := config.AppPort()
	clientURL := config.ClientURL()
	isProd := config.IsProd()

	db := adapter.ConnectSQL(connection)
	validations.TryRegister(binding.Validator.Engine())
	router, server := adapter.CreateRouter(port, isProd, clientURL)
	api := adapter.CreateRoutes(router, logger)
	modules.Init(api, db)

	api.GET("/test", func(context *gin.Context) {
		for i := 0; i <= 30; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}

		context.JSON(http.StatusOK, gin.H{"message": "Test passed"})
	})

	adapter.ServeGracefully(server, logger, SHUTDOWN_TIMEOUT)
}

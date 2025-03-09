package main

import (
	"time"

	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/common/config"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"

	validator "github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/adapters/playground"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	log := log.New("Config")
	log.Info("Parsed environment variables")

	connection := config.DBConnection()
	db, _ := server.ConnectSQL(connection, 1*time.Second)
	db.Close()
	log.Info("Connected to database")

	validator.TryRegister(binding.Validator.Engine())
	log.Info("Registered validations")
}

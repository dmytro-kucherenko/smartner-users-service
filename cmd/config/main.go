package main

import (
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/common/config"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/validations"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	log := log.New("Config")
	log.Info("Parsed environment variables")

	connection := config.DBConnection()
	db := adapter.ConnectSQL(connection)
	db.Close()
	log.Info("Connected to database")

	validations.TryRegister(binding.Validator.Engine())
	log.Info("Registered validations")
}

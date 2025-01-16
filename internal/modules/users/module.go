package users

import (
	"database/sql"

	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules/users/controllers"
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules/users/repositories"
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules/users/services"
	"github.com/gin-gonic/gin"
)

type Module struct {
	repository *repositories.Main
	service    *services.Main
	controller *controllers.Main
}

func NewModule(connection *sql.DB) *Module {
	repository := repositories.New(connection)
	service := services.New(repository)
	controller := controllers.New(service)

	return &Module{repository, service, controller}
}

func (module *Module) Init(router *gin.RouterGroup) {
	module.controller.Init(router)
}

func (module *Module) Service(router *gin.RouterGroup) *services.Main {
	return module.service
}

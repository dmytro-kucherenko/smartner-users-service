package user

import (
	"database/sql"

	"github.com/dmytro-kucherenko/smartner-contracts-package/pkg/modules/user"
	"github.com/dmytro-kucherenko/smartner-users-service/internal/modules/user/repositories"
	"github.com/dmytro-kucherenko/smartner-users-service/internal/modules/user/services"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	adapterGin "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
	adapterGRPC "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/grpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Module struct {
	repository *repositories.Main
	service    *services.Main
	controller *Controller
	caller     *Caller
}

func NewModule(db *sql.DB, userConn *grpc.ClientConn) *Module {
	userClient := user.NewClient(userConn)
	repository := repositories.New(db)
	service := services.New(repository, userClient)
	controller := NewController(service)
	caler := NewCaller(service)

	return &Module{repository, service, controller, caler}
}

func (module *Module) Service(router *gin.RouterGroup) *services.Main {
	return module.service
}

func (module *Module) Controllers() []adapterGin.Controller {
	return []adapterGin.Controller{module.controller}
}

func (module *Module) Callers() []adapterGRPC.Caller {
	return []adapterGRPC.Caller{module.caller}
}

func (module *Module) Modules() []server.Module {
	return []server.Module{}
}

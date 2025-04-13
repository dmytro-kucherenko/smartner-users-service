package modules

import (
	"database/sql"

	"github.com/dmytro-kucherenko/smartner-users-service/internal/modules/user"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"google.golang.org/grpc"
)

type App struct {
	userModule *user.Module
}

func NewApp(db *sql.DB, userConn *grpc.ClientConn) *App {
	userModule := user.NewModule(db, userConn)

	return &App{userModule}
}

func (app *App) Modules() []server.Module {
	return []server.Module{app.userModule}
}

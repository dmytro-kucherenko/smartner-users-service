package modules

import (
	"database/sql"

	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules/users"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/gin-gonic/gin"
)

func Init(api *gin.RouterGroup, db *sql.DB, meta server.RequestMeta) {
	usersModule := users.NewModule(db)
	usersModule.Init(api, meta)
}

package modules

import (
	"database/sql"

	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules/users"
	"github.com/gin-gonic/gin"
)

func Init(api *gin.RouterGroup, db *sql.DB) {
	usersModule := users.NewModule(db)
	usersModule.Init(api)
}

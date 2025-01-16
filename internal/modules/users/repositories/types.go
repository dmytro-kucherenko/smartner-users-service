package repositories

import (
	db "github.com/Dmytro-Kucherenko/smartner-users-service/internal/gen/db/users"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/pagination"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
)

type FindOneParams struct {
	ID    types.Optional[types.ID]
	Email types.Optional[string]
}

type FindPageParams struct {
	pagination.PageMeasures
}

type CreateParams struct {
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	PasswordSalt string
}

type UpdateParams struct {
	FirstName    types.Optional[string]
	LastName     types.Optional[string]
	PasswordHash types.Optional[string]
	PasswordSalt types.Optional[string]
}

type ItemQuery = db.User

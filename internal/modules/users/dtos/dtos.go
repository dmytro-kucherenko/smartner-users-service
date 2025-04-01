package dtos

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/pagination/dtos"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
)

type GetParams struct {
	ID types.ID `uri:"id" validate:"required,uuid4" swaggertype:"string"`
} // @name UserGetParamsDto

type GetAllParams struct {
	dtos.PageParams
} // @name UsersGetAllParamsDto

type SignInParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
} // @name UserSignInParamsDto

type SignUpParams struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password"`
} // @name UserSignUpParamsDto

type UpdateParams struct {
	ID        types.ID `uri:"id" validate:"required,uuid4" swaggertype:"string"`
	FirstName *string  `json:"firstName"`
	LastName  *string  `json:"lastName"`
} // @name UserUpdateParamsDto

type Item struct {
	ID        types.ID `json:"id" swaggertype:"string"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
} // @name UserItemDto

type Page struct {
	Items []Item        `json:"items"`
	Meta  dtos.PageMeta `json:"meta"`
} // @name UsersPageDto

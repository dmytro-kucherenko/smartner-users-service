package dtos

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/pagination/dtos"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
)

type GetRequest struct {
	ID types.ID `uri:"id" binding:"required" swaggertype:"string"`
} // @name UserGetParamsDto

type GetAllRequest struct {
	dtos.PageQueryRequest
} // @name UsersGetAllQueryDto

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
} // @name UserSignInBodyDto

type SignUpRequest struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required,password"`
} // @name UserSignUpBodyDto

type UpdateRequest struct {
	FirstName *string `json:"firstname"`
	LastName  *string `json:"lastname"`
} // @name UserUpdateBodyDto

type ItemResponse struct {
	ID        types.ID `json:"id" swaggertype:"string"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"secondname"`
} // @name UserItemDto

type PageResponse struct {
	Items []ItemResponse        `json:"items"`
	Meta  dtos.PageMetaResponse `json:"meta"`
} // @name UsersPageDto

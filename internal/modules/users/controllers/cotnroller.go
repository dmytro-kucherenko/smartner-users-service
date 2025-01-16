package controllers

import (
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules/users/dtos"
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules/users/services"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
	"github.com/gin-gonic/gin"
)

type Main struct {
	service *services.Main
}

func New(service *services.Main) *Main {
	return &Main{service}
}

func (controller *Main) Init(router *gin.RouterGroup) {
	group := router.Group("users")
	adapter.Get(group, adapter.Config("/get/:id", 200, nil), controller.get)
	adapter.Get(group, adapter.Config("/page", 200, nil), controller.getPage)
	adapter.Post(group, adapter.Config("/signIn", 200, nil), controller.signIn)
	adapter.Post(group, adapter.Config("/signUp", 201, nil), controller.signUp)
	adapter.Put(group, adapter.Config("/update/:id", 200, nil), controller.update)
	adapter.Delete(group, adapter.Config("/delete/:id", 200, nil), controller.delete)
}

// @Summary	Get user
// @Tags		Users
// @Accept		json
// @Produce	json
// @Param		path	path		UserGetParamsDto	true	"User Filters"
// @Success	200		{object}	UserItemDto
// @Failure	400		{object}	ErrorDto
// @Failure	404		{object}	ErrorDto
// @Router		/users/get/:id [get]
func (controller *Main) get(options *server.RequestOptions[any, dtos.GetRequest, any]) (dtos.ItemResponse, error) {
	return controller.service.Get(options.Ctx, options.Params)
}

// @Summary	Get users page
// @Tags		Users
// @Accept		json
// @Produce	json
// @Param		name	query		UsersGetAllQueryDto	true	"Page Filters"
// @Success	200		{object}	UsersPageDto
// @Failure	400		{object}	ErrorDto
// @Router		/users/page [get]
func (controller *Main) getPage(options *server.RequestOptions[any, any, dtos.GetAllRequest]) (dtos.PageResponse, error) {
	return controller.service.GetPage(options.Ctx, options.Query)
}

// @Summary	Sign in user
// @Tags		Users
// @Accept		json
// @Produce	json
// @Param		body	body		UserSignInBodyDto	true	"User Data"
// @Success	200		{object}	UserItemDto
// @Failure	400		{object}	ErrorDto
// @Failure	401		{object}	ErrorDto
// @Failure	409		{object}	ErrorDto
// @Router		/users/signIn [post]
func (controller *Main) signIn(options *server.RequestOptions[dtos.SignInRequest, any, any]) (dtos.ItemResponse, error) {
	return controller.service.SignIn(options.Ctx, options.Body)
}

// @Summary	Sign up user
// @Tags		Users
// @Accept		json
// @Produce	json
// @Param		body	body		UserSignUpBodyDto	true	"User Data"
// @Success	200		{object}	UserItemDto
// @Failure	400		{object}	ErrorDto
// @Failure	401		{object}	ErrorDto
// @Failure	409		{object}	ErrorDto
// @Router		/users/signUp [post]
func (controller *Main) signUp(options *server.RequestOptions[dtos.SignUpRequest, any, any]) (dtos.ItemResponse, error) {
	return controller.service.SignUp(options.Ctx, options.Body)
}

// @Summary	Update user
// @Tags		Users
// @Security JWTAuth
// @Accept		json
// @Produce	json
// @Param		body	body		UserUpdateBodyDto	true	"User Data"
// @Success	200		{object}	UserItemDto
// @Failure	400		{object}	ErrorDto
// @Failure	401		{object}	ErrorDto
// @Failure	404		{object}	ErrorDto
// @Router		/users/update/:id [put]
func (controller *Main) update(options *server.RequestOptions[dtos.UpdateRequest, dtos.GetRequest, any]) (dtos.ItemResponse, error) {
	return controller.service.Update(options.Ctx, options.Params, options.Body)
}

// @Summary	Delete user
// @Tags		Users
// @Security JWTAuth
// @Accept		json
// @Produce	json
// @Param		path	path		UserGetParamsDto	true	"User Filters"
// @Success	200		{object}	UserItemDto
// @Failure	400		{object}	ErrorDto
// @Failure	401		{object}	ErrorDto
// @Failure	404		{object}	ErrorDto
// @Router		/users/delete/:id [delete]
func (controller *Main) delete(options *server.RequestOptions[any, dtos.GetRequest, any]) (dtos.ItemResponse, error) {
	return controller.service.Delete(options.Ctx, options.Params)
}

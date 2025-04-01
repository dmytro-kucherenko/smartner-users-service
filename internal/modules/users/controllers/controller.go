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

func (controller *Main) Init(group *gin.RouterGroup, meta server.RequestMeta) {
	config := adapter.NewConfig(meta).WithSession()
	adapter.Get(group, controller.get, config.MapRoute("/get/:id", 200))
	adapter.Get(group, controller.getPage, config.MapRoute("/page", 200))
	adapter.Post(group, controller.signIn, config.MapRoute("/signIn", 200))
	adapter.Post(group, controller.signUp, config.MapRoute("/signUp", 201))
	adapter.Put(group, controller.update, config.MapRoute("/update/:id", 200))
	adapter.Delete(group, controller.delete, config.MapRoute("/delete/:id", 200))
}

// @Summary	Get user
// @Tags		Users
// @Security JWTAuth
// @Accept		json
// @Produce	json
// @Param		path	path		UserGetParamsDto	true	"User Filters"
// @Success	200		{object}	UserItemDto
// @Failure	400		{object}	ErrorDto
// @Failure	404		{object}	ErrorDto
// @Router		/users/get/:id [get]
func (controller *Main) get(options *server.RequestOptions[dtos.GetParams]) (dtos.Item, error) {
	return controller.service.Get(options.Ctx, options.Params)
}

// @Summary	Get users page
// @Tags		Users
// @Security JWTAuth
// @Accept		json
// @Produce	json
// @Param		name	query		UserGetAllParamsDto	true	"Page Filters"
// @Success	200		{object}	UserPageDto
// @Failure	400		{object}	ErrorDto
// @Router		/users/page [get]
func (controller *Main) getPage(options *server.RequestOptions[dtos.GetAllParams]) (dtos.Page, error) {
	return controller.service.GetPage(options.Ctx, options.Params)
}

// @Summary	Sign in user
// @Tags		Users
// @Accept		json
// @Produce	json
// @Param		body	body		UserSignInParamsDto	true	"User Data"
// @Success	200		{object}	UserItemDto
// @Failure	400		{object}	ErrorDto
// @Failure	401		{object}	ErrorDto
// @Failure	409		{object}	ErrorDto
// @Router		/users/signIn [post]
func (controller *Main) signIn(options *server.RequestOptions[dtos.SignInParams]) (dtos.Item, error) {
	return controller.service.SignIn(options.Ctx, options.Params)
}

// @Summary	Sign up user
// @Tags		Users
// @Security JWTAuth
// @Accept		json
// @Produce	json
// @Param		body	body		UserSignUpParamsDto	true	"User Data"
// @Success	200		{object}	UserItemDto
// @Failure	400		{object}	ErrorDto
// @Failure	401		{object}	ErrorDto
// @Failure	409		{object}	ErrorDto
// @Router		/users/signUp [post]
func (controller *Main) signUp(options *server.RequestOptions[dtos.SignUpParams]) (dtos.Item, error) {
	return controller.service.SignUp(options.Ctx, options.Params)
}

// @Summary	Update user
// @Tags		Users
// @Security JWTAuth
// @Accept		json
// @Produce	json
// @Param		body	body		UserUpdateParamsDto	true	"User Data"
// @Success	200		{object}	UserItemDto
// @Failure	400		{object}	ErrorDto
// @Failure	401		{object}	ErrorDto
// @Failure	404		{object}	ErrorDto
// @Router		/users/update/:id [put]
func (controller *Main) update(options *server.RequestOptions[dtos.UpdateParams]) (dtos.Item, error) {
	return controller.service.Update(options.Ctx, options.Params)
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
func (controller *Main) delete(options *server.RequestOptions[dtos.GetParams]) (dtos.Item, error) {
	return controller.service.Delete(options.Ctx, options.Params)
}

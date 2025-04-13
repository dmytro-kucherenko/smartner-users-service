package user

import (
	"context"
	"net/http"

	"github.com/dmytro-kucherenko/smartner-users-service/internal/modules/user/dtos"
	"github.com/dmytro-kucherenko/smartner-users-service/internal/modules/user/services"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin/interceptors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *services.Main
}

func NewController(service *services.Main) *Controller {
	return &Controller{service}
}

func (controller *Controller) Init(group *gin.RouterGroup) {
	config := adapter.NewConfig().WithInterceptor(interceptors.Auth())
	adapter.Get(group, controller.get, config.MapRoute("/get/:id", http.StatusOK))
	adapter.Get(group, controller.getPage, config.MapRoute("/page", http.StatusOK))
	adapter.Post(group, controller.signIn, config.MapRoute("/signIn", http.StatusOK))
	adapter.Post(group, controller.signUp, config.MapRoute("/signUp", http.StatusCreated))
	adapter.Put(group, controller.update, config.MapRoute("/update/:id", http.StatusOK))
	adapter.Delete(group, controller.delete, config.MapRoute("/delete/:id", http.StatusOK))
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
func (controller *Controller) get(ctx context.Context, params dtos.GetParamsDTO) (dtos.Item, error) {
	return controller.service.Get(ctx, params)
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
func (controller *Controller) getPage(ctx context.Context, params dtos.GetAllParams) (dtos.Page, error) {
	return controller.service.GetPage(ctx, params)
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
func (controller *Controller) signIn(ctx context.Context, params dtos.SignInParams) (dtos.Item, error) {
	return controller.service.SignIn(ctx, params)
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
func (controller *Controller) signUp(ctx context.Context, params dtos.SignUpParams) (dtos.Item, error) {
	return controller.service.SignUp(ctx, params)
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
func (controller *Controller) update(ctx context.Context, params dtos.UpdateParams) (dtos.Item, error) {
	return controller.service.Update(ctx, params)
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
func (controller *Controller) delete(ctx context.Context, params dtos.GetParamsDTO) (dtos.Item, error) {
	return controller.service.Delete(ctx, params)
}

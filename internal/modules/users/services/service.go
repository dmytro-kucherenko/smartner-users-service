package services

import (
	"context"
	"net/http"

	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/common/config"
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules/users/dtos"
	"github.com/Dmytro-Kucherenko/smartner-users-service/internal/modules/users/repositories"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/encrypt"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/pagination"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
)

type Main struct {
	repository     *repositories.Main
	encryptService *encrypt.Service
}

func New(repository *repositories.Main) *Main {
	return &Main{
		repository:     repository,
		encryptService: encrypt.NewService(config.PasswordSecret(), config.PasswordRounds()),
	}
}

func (service *Main) Get(ctx context.Context, filters dtos.GetParams) (user dtos.Item, err error) {
	item, err := service.repository.FindOne(ctx, repositories.FindOneParams{
		ID: types.OptionalValue(filters.ID),
	})

	if err != nil {
		err = errors.NewHttpError(http.StatusNotFound, "user was not found.")

		return
	}

	user = transform(item)[0]

	return
}

func (service *Main) GetPage(ctx context.Context, params dtos.GetAllParams) (page dtos.Page, err error) {
	total, err := service.repository.Count(ctx)
	if err != nil {
		return
	}

	measures, err := pagination.GetPageMeasures(total, params.PageParams)
	if err != nil {
		return
	}

	items, err := service.repository.FindPage(ctx, repositories.FindPageParams{PageMeasures: measures})
	if err != nil {
		return
	}

	page = dtos.Page{
		Items: transform(items...),
		Meta:  pagination.GetPageMeta(total, measures),
	}

	return
}

func (service *Main) SignIn(ctx context.Context, params dtos.SignInParams) (user dtos.Item, err error) {
	item, err := service.repository.FindOne(ctx, repositories.FindOneParams{
		Email: types.OptionalValue(params.Email),
	})

	if err != nil {
		err = errors.NewHttpError(http.StatusUnauthorized, "user was not found.")

		return
	}

	ok := service.encryptService.Verify(params.Password, encrypt.Value{
		Hash: item.PasswordHash,
		Salt: item.PasswordSalt,
	})

	if !ok {
		err = errors.NewHttpError(http.StatusUnauthorized, "user password is not correct.")

		return
	}

	user = transform(item)[0]

	return
}

func (service *Main) SignUp(ctx context.Context, params dtos.SignUpParams) (user dtos.Item, err error) {
	_, err = service.repository.FindOne(ctx, repositories.FindOneParams{
		Email: types.OptionalValue(params.Email),
	})

	if err == nil {
		err = errors.NewHttpError(http.StatusConflict, "user with this email already exists.")

		return
	}

	password, err := service.encryptService.Gen(params.Password)
	if err != nil {
		return
	}

	item, err := service.repository.Create(ctx, repositories.CreateParams{
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		Email:        params.Email,
		PasswordHash: password.Hash,
		PasswordSalt: password.Salt,
	})

	if err != nil {
		return
	}

	user = transform(item)[0]

	return
}

// Separate route to update password with previous one
func (service *Main) Update(ctx context.Context, params dtos.UpdateParams) (user dtos.Item, err error) {
	user, err = service.Get(ctx, dtos.GetParams{ID: params.ID})
	if err != nil {
		return
	}

	item, err := service.repository.Update(ctx, params.ID, &repositories.UpdateParams{
		FirstName: types.OptionalPointer(params.FirstName),
		LastName:  types.OptionalPointer(params.LastName),
	})

	if err != nil {
		return
	}

	user = transform(item)[0]

	return
}

func (service *Main) Delete(ctx context.Context, filters dtos.GetParams) (user dtos.Item, err error) {
	_, err = service.Get(ctx, filters)
	if err != nil {
		return
	}

	item, err := service.repository.Delete(ctx, filters.ID)
	if err != nil {
		return
	}

	user = transform(item)[0]

	return
}

func transform(items ...repositories.ItemQuery) []dtos.Item {
	users := make([]dtos.Item, 0, len(items))

	for _, item := range items {
		users = append(users, dtos.Item{
			ID:        item.ID,
			FirstName: item.FirstName,
			LastName:  item.LastName,
		})
	}

	return users
}

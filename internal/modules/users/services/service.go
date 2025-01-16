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

func (service *Main) Get(ctx context.Context, filters dtos.GetRequest) (user dtos.ItemResponse, err error) {
	item, err := service.repository.FindOne(ctx, repositories.FindOneParams{
		ID: types.OptionalValue(filters.ID),
	})

	if err != nil {
		err = errors.NewHttpError(http.StatusNotFound, "User was not found.")

		return
	}

	user = transform(item)[0]

	return
}

func (service *Main) GetPage(ctx context.Context, payload dtos.GetAllRequest) (page dtos.PageResponse, err error) {
	total, err := service.repository.Count(ctx)
	if err != nil {
		return
	}

	measures, err := pagination.GetPageMeasures(total, payload.PageQueryRequest)
	if err != nil {
		return
	}

	items, err := service.repository.FindPage(ctx, repositories.FindPageParams{PageMeasures: measures})
	if err != nil {
		return
	}

	page = dtos.PageResponse{
		Items: transform(items...),
		Meta:  pagination.GetPageMeta(total, measures),
	}

	return
}

func (service *Main) SignIn(ctx context.Context, payload dtos.SignInRequest) (user dtos.ItemResponse, err error) {
	item, err := service.repository.FindOne(ctx, repositories.FindOneParams{
		Email: types.OptionalValue(payload.Email),
	})

	if err != nil {
		err = errors.NewHttpError(http.StatusUnauthorized, "User was not found.")

		return
	}

	ok := service.encryptService.Verify(payload.Password, encrypt.Value{
		Hash: item.PasswordHash,
		Salt: item.PasswordSalt,
	})

	if !ok {
		err = errors.NewHttpError(http.StatusUnauthorized, "User password is not correct.")

		return
	}

	user = transform(item)[0]

	return
}

func (service *Main) SignUp(ctx context.Context, payload dtos.SignUpRequest) (user dtos.ItemResponse, err error) {
	_, err = service.repository.FindOne(ctx, repositories.FindOneParams{
		Email: types.OptionalValue(payload.Email),
	})

	if err == nil {
		err = errors.NewHttpError(http.StatusConflict, "User with this email already exists.")

		return
	}

	password, err := service.encryptService.Gen(payload.Password)
	if err != nil {
		return
	}

	item, err := service.repository.Create(ctx, repositories.CreateParams{
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
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
func (service *Main) Update(ctx context.Context, filters dtos.GetRequest, payload dtos.UpdateRequest) (user dtos.ItemResponse, err error) {
	user, err = service.Get(ctx, filters)
	if err != nil {
		return
	}

	item, err := service.repository.Update(ctx, filters.ID, &repositories.UpdateParams{
		FirstName: types.OptionalPointer(payload.FirstName),
		LastName:  types.OptionalPointer(payload.LastName),
	})

	if err != nil {
		return
	}

	user = transform(item)[0]

	return
}

func (service *Main) Delete(ctx context.Context, filters dtos.GetRequest) (user dtos.ItemResponse, err error) {
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

func transform(items ...repositories.ItemQuery) []dtos.ItemResponse {
	users := make([]dtos.ItemResponse, 0, len(items))

	for _, item := range items {
		users = append(users, dtos.ItemResponse{
			ID:        item.ID,
			FirstName: item.FirstName,
			LastName:  item.LastName,
		})
	}

	return users
}

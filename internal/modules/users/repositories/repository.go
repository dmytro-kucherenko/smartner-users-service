package repositories

import (
	"context"
	"database/sql"

	db "github.com/Dmytro-Kucherenko/smartner-users-service/internal/gen/db/users"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/queries"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
	"github.com/google/uuid"
)

type Main struct {
	manager queries.Manager[*db.Queries]
	queries.TransactionCaller[*db.Queries]
}

func New(connection *sql.DB) *Main {
	queries := queries.NewSQLManager(
		connection,
		func(connection *sql.DB) *db.Queries { return db.New(connection) },
		func(ctx context.Context, connection *sql.DB) (*db.Queries, error) { return db.Prepare(ctx, connection) },
	)

	return &Main{
		manager:           queries,
		TransactionCaller: queries,
	}
}

func (service *Main) queries() *db.Queries {
	return service.manager.Queries()
}

func (service *Main) Count(ctx context.Context) (total uint64, err error) {
	count, err := service.queries().Count(ctx)
	if err != nil {
		return
	}

	return uint64(count), nil
}

func (service *Main) FindOne(ctx context.Context, filters FindOneParams) (ItemQuery, error) {
	return service.queries().FindOne(ctx, db.FindOneParams{
		ID:    uuid.NullUUID{UUID: uuid.UUID(filters.ID.Value), Valid: filters.ID.Valid},
		Email: sql.NullString{String: filters.Email.Value, Valid: filters.Email.Valid},
	})
}

func (service *Main) FindPage(ctx context.Context, filters FindPageParams) ([]ItemQuery, error) {
	items, err := service.queries().FindPage(ctx, db.FindPageParams{
		Offset: int32(filters.Offset),
		Limit:  int32(filters.Limit),
	})

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (service *Main) Create(ctx context.Context, payload CreateParams) (ItemQuery, error) {
	user, err := service.queries().Create(ctx, db.CreateParams{
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		PasswordHash: payload.PasswordHash,
		PasswordSalt: payload.PasswordSalt,
	})

	return user, err
}

func (service *Main) Update(ctx context.Context, ID types.ID, payload *UpdateParams) (ItemQuery, error) {
	return service.queries().Update(ctx, db.UpdateParams{
		ID:           ID,
		FirstName:    sql.NullString{String: payload.FirstName.Value, Valid: payload.FirstName.Valid},
		LastName:     sql.NullString{String: payload.LastName.Value, Valid: payload.LastName.Valid},
		PasswordHash: sql.NullString{String: payload.PasswordHash.Value, Valid: payload.PasswordHash.Valid},
		PasswordSalt: sql.NullString{String: payload.PasswordSalt.Value, Valid: payload.PasswordSalt.Valid},
	})
}

func (service *Main) Delete(ctx context.Context, ID types.ID) (ItemQuery, error) {
	return service.queries().Delete(ctx, ID)
}

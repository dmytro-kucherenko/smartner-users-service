// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
)

type User struct {
	ID           types.ID
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	PasswordSalt string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

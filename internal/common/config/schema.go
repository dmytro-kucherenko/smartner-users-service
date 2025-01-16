package config

import (
	helpers "github.com/dmytro-kucherenko/smartner-utils-package/pkg/config"
)

type Schema struct {
	AppEnv         string `validate:"required,oneof=local stage prod"`
	AppPort        int    `validate:"required"`
	ClientURL      string `validate:"required"`
	DBHost         string `validate:"required"`
	DBPort         int    `validate:"required"`
	DBName         string `validate:"required"`
	DBUsername     string `validate:"required"`
	DBPassword     string `validate:"required"`
	DBConnection   string `validate:"required"`
	PasswordSecret string `validate:"required"`
	PasswordRounds uint8  `validate:"required,min=4"`
}

var schema Schema

func init() {
	schema = helpers.Init(".env", func() Schema {
		return Schema{
			AppEnv:         helpers.GetEnvString("APP_ENV"),
			AppPort:        helpers.GetEnvInt("APP_PORT"),
			ClientURL:      helpers.GetEnvString("CLIENT_URL"),
			DBHost:         helpers.GetEnvString("DB_HOST"),
			DBPort:         helpers.GetEnvInt("DB_PORT"),
			DBName:         helpers.GetEnvString("DB_NAME"),
			DBUsername:     helpers.GetEnvString("DB_USERNAME"),
			DBPassword:     helpers.GetEnvString("DB_PASSWORD"),
			DBConnection:   helpers.GetEnvString("DB_CONNECTION"),
			PasswordSecret: helpers.GetEnvString("PASSWORD_SECRET"),
			PasswordRounds: uint8(helpers.GetEnvInt("PASSWORD_ROUNDS")),
		}
	})
}

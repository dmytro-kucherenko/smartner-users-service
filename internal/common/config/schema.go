package config

import (
	helpers "github.com/dmytro-kucherenko/smartner-utils-package/pkg/config"
)

func GetEnvBool(key string) bool {
	value := helpers.GetEnvInt(key)

	return value != 0
}

type Schema struct {
	AppEnv         string `validate:"required,oneof=local stage prod"`
	AppPort        uint16
	AppProtocol    string `validate:"required,oneof=http https"`
	AppHost        string `validate:"required"`
	AppBasePath    string
	AppOnlyConfig  bool   `validate:"required"`
	ClientURL      string `validate:"required"`
	DBHost         string `validate:"required"`
	DBPort         uint16 `validate:"required"`
	DBName         string `validate:"required"`
	DBUsername     string `validate:"required"`
	DBPassword     string `validate:"required"`
	DBSchema       string `validate:"required"`
	PasswordSecret string `validate:"required"`
	PasswordRounds uint8  `validate:"required,min=4"`
	Test           string `validate:"required"`
}

var schema Schema

func Load() (err error) {
	schema, err = helpers.Init(".env", func() Schema {
		return Schema{
			AppEnv:         helpers.GetEnvString("APP_ENV"),
			AppPort:        uint16(helpers.GetEnvInt("APP_PORT")),
			AppProtocol:    helpers.GetEnvString("APP_PROTOCOL"),
			AppHost:        helpers.GetEnvString("APP_HOST"),
			AppBasePath:    helpers.GetEnvString("APP_BASE_PATH"),
			AppOnlyConfig:  GetEnvBool("APP_ONLY_CONFIG"),
			ClientURL:      helpers.GetEnvString("CLIENT_URL"),
			DBHost:         helpers.GetEnvString("DB_HOST"),
			DBPort:         uint16(helpers.GetEnvInt("DB_PORT")),
			DBName:         helpers.GetEnvString("DB_NAME"),
			DBUsername:     helpers.GetEnvString("DB_USERNAME"),
			DBPassword:     helpers.GetEnvString("DB_PASSWORD"),
			DBSchema:       helpers.GetEnvString("DB_SCHEMA"),
			PasswordSecret: helpers.GetEnvString("PASSWORD_SECRET"),
			PasswordRounds: uint8(helpers.GetEnvInt("PASSWORD_ROUNDS")),
		}
	})

	return
}

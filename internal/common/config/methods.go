package config

import (
	"fmt"
	"strings"
)

const PortDefault uint16 = 8000

func IsProd() bool {
	return schema.AppEnv == "prod"
}

func IsLocal() bool {
	return schema.AppEnv == "local" || schema.AppEnv == "stage"
}

func AppEnv() string {
	return schema.AppEnv
}

func AppPort() uint16 {
	if schema.AppPort == 0 {
		return PortDefault
	}

	return schema.AppPort
}

func AppProtocol() string {
	return schema.AppProtocol
}

func AppHost() string {
	return schema.AppHost
}

func AppBasePath() string {
	if schema.AppBasePath != "" && !strings.HasPrefix(schema.AppBasePath, "/") {
		return "/" + schema.AppBasePath
	}

	return schema.AppBasePath
}

func AppURL() string {
	return fmt.Sprintf("%v://%v%v", AppProtocol(), AppHost(), AppBasePath())
}

func AppOnlyConfig() bool {
	return schema.AppOnlyConfig
}

func ClientURL() string {
	return schema.ClientURL
}

func DBHost() string {
	return schema.DBHost
}

func DBPort() uint16 {
	return schema.AppPort
}

func DBName() string {
	return schema.DBName
}

func DBUsername() string {
	return schema.DBUsername
}

func DBPassword() string {
	return schema.DBPassword
}

func DBConnection() string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?search_path=%v",
		schema.DBUsername,
		schema.DBPassword,
		schema.DBHost,
		schema.DBPort,
		schema.DBName,
		schema.DBSchema,
	)
}

func PasswordSecret() string {
	return schema.PasswordSecret
}

func PasswordRounds() uint8 {
	return schema.PasswordRounds
}

package config

func IsProd() bool {
	return schema.AppEnv == "prod"
}

func IsLocal() bool {
	return schema.AppEnv == "local"
}

func AppEnv() string {
	return schema.AppEnv
}

func AppPort() int {
	return schema.AppPort
}

func ClientURL() string {
	return schema.ClientURL
}

func DBHost() string {
	return schema.DBHost
}

func DBPort() int {
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
	return schema.DBConnection
}

func PasswordSecret() string {
	return schema.PasswordSecret
}

func PasswordRounds() uint8 {
	return schema.PasswordRounds
}

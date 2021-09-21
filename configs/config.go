package configs

import (
	"errors"
	"os"
)

type Config interface {
	GetPort() string
	GetDatabaseHost() string
	GetDatabaseName() string
	GetDatabasePort() string
	GetDatabaseUser() string
	GetDatabasePass() string
}

type ConfigImpl struct {
	port         string
	databaseHost string
	databaseName string
	databasePort string
	databaseUser string
	databasePass string
}

func NewConfig() Config {
	return &ConfigImpl{
		port:         getValue("PORT"),
		databaseHost: getValue("DATABASE_HOST"),
		databaseName: getValue("DATABASE_NAME"),
		databasePort: getValue("DATABASE_PORT"),
		databaseUser: getValue("DATABASE_USER"),
		databasePass: getValue("DATABASE_PASS"),
	}
}

func getValue(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(errors.New(key + " doesn't exist."))
	}

	return value
}

func (c ConfigImpl) GetPort() string {
	return c.port
}

func (c ConfigImpl) GetDatabaseHost() string {
	return c.databaseHost
}

func (c ConfigImpl) GetDatabaseName() string {
	return c.databaseName
}

func (c ConfigImpl) GetDatabasePort() string {
	return c.databasePort
}

func (c ConfigImpl) GetDatabaseUser() string {
	return c.databaseUser
}

func (c ConfigImpl) GetDatabasePass() string {
	return c.databasePass
}

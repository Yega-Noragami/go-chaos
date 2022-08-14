package config

import (
	"os"

	"github.com/Yega-Noragami/go-chaos/pkg/constants"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUser     string
}

func NewConfig() *Config {
	configStore := Config{}
	val, ok := os.LookupEnv(constants.DBHostKey)
	if !ok || val == "" {
		val = constants.SQL_HOSTNAME_LOCAL
	}
	configStore.DBHost = val
	val, ok = os.LookupEnv(constants.DBPortKey)
	if !ok || val == "" {
		val = constants.SQL_PORT_LOCAL
	}
	configStore.DBPort = val
	val, ok = os.LookupEnv(constants.DBDatabaseKey)
	if !ok || val == "" {
		val = constants.SQL_DB_LOCAL
	}
	configStore.DBDatabase = val
	val, ok = os.LookupEnv(constants.DBUserKey)
	if !ok || val == "" {
		val = constants.SQL_USER_LOCAL
	}
	configStore.DBUser = val
	return &configStore
}

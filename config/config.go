package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	App         string
	Environment string
	Version     string

	ServiceHost string
	HTTPPort    string

	PostgresHost            string
	PostgresPort            int
	PostgresUser            string
	PostgresPassword        string
	PostgresDatabase        string
	PostgresMaxConnections  int
	PostgresConnMaxIdleTime int
	DefaultOffset           string
	DefaultLimit            string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No.. .env file ")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("PROJECT_NAME", "JASURBEK"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "dev"))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.PostgresMaxConnections = cast.ToInt(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 5))
	config.PostgresConnMaxIdleTime = cast.ToInt(getOrReturnDefaultValue("POSTGRES_CONN_MAX_IDLE_TIME", 5))

	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":9090"))
	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "secret_password"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DB", "postgres"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
package config

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	AppEnv              string
	DbSource            string
	JwtSecret           []byte
	JwtExpired          string
	HttpAddress         string
	OpenapiYamlLocation string
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		AppEnv:              getEnv("APP_ENV", "production"),
		DbSource:            getEnv("DB_SOURCE", "dashboard.db"),
		JwtSecret:           []byte(getEnv("JWT_SECRET", "dev-secret-replace-me")),
		JwtExpired:          getEnv("JWT_EXPIRED", "24h"),
		HttpAddress:         getEnv("HTTP_ADDR", ":8080"),
		OpenapiYamlLocation: getEnv("OPENAPIYAML_LOCATION", "../openapi.yaml"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

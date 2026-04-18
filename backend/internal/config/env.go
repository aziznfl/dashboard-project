package config

import (
	"os"
	"strconv"

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
	RedisHost           string
	RedisPort           string
	RedisPassword       string
	RedisDB             int
}

func Load() *Config {
	_ = godotenv.Load()
	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	return &Config{
		AppEnv:              getEnv("APP_ENV", "production"),
		DbSource:            getEnv("DB_SOURCE", "dashboard.db"),
		JwtSecret:           []byte(getEnv("JWT_SECRET", "dev-secret-replace-me")),
		JwtExpired:          getEnv("JWT_EXPIRED", "24h"),
		HttpAddress:         getEnv("HTTP_ADDR", ":8080"),
		OpenapiYamlLocation: getEnv("OPENAPIYAML_LOCATION", "../openapi.yaml"),
		RedisHost:           getEnv("REDIS_HOST", "localhost"),
		RedisPort:           getEnv("REDIS_PORT", "6379"),
		RedisPassword:       getEnv("REDIS_PASSWORD", ""),
		RedisDB:             redisDB,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

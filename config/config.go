package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No se encontro archivo .env")
	}

	appConfig := &Config{
		Port:        getEnv("PORT", ":5050"),
		JWTSecret:   getEnv("JWT_SECRET", "mysecretkey"),
		DatabaseURL: getEnv("DATABASE_URL", "mysql://root:@tcp(localhost:3306)/dbname"),
	}
	return appConfig
}

func getEnv(key string, defatulValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defatulValue
	}
	return value
}

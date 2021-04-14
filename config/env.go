package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Configuration struct {
	Debug    bool
	Key      string
	Jwt      Jwt
	Database Database
}

var Config *Configuration

func Load() *Configuration {
	if Config == nil {
		_ = godotenv.Load()
		Config = &Configuration{}
		Config.Key = os.Getenv("APP_KEY")
		Config.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
		Config.Jwt.LifeTime, _ = strconv.Atoi(os.Getenv("JWT_LIFETIME"))
		Config.Database = Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		}
	}
	return Config
}

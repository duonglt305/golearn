package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Configuration struct {
	Debug bool
}

var Config *Configuration

func Load() *Configuration {
	if Config == nil {
		_ = godotenv.Load()
		Config = &Configuration{}
		Config.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	}
	return Config
}

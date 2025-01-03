package helpers

import (
	"github.com/joho/godotenv"
)

var (
	EnvMap map[string]string
)

func SetupConfig() {
	var err error
	EnvMap, err = godotenv.Read(".env")
	if err != nil {
		Logger.Fatal("Error loading .env file: ", err)
	}
	Logger.Info("Config loaded")
}

func GetEnv(key string) string {
	return EnvMap[key]
}

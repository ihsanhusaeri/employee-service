package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//List of db configuration name
var (
	AppPort    string
	DbUsername string
	DbPassword string
	DbName     string
	URLDb      string
	DbPort     string
)

func getEnv(key string, fileEnv map[string]string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return fileEnv[key]
	}
	return envVal
}

//InitEnv to read env file based on chosen mode
func InitEnv() {
	envMode := os.Getenv("NODE_ENV")

	if envMode == "" {
		envMode = "development"
	}

	envFile, err := godotenv.Read("env/" + envMode + ".env")

	if err != nil {
		fmt.Printf("Error read %s.env", envMode)
	}

	AppPort = getEnv("APP_PORT", envFile)
	DbUsername = getEnv("DB_USERNAME", envFile)
	DbPassword = getEnv("DB_PASSWORD", envFile)
	DbName = getEnv("DB_NAME", envFile)
	URLDb = getEnv("URL_DB", envFile)
	DbPort = getEnv("DB_PORT", envFile)

}

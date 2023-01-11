package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warn:", err)
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return ""
}

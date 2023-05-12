package infrastructure

import (
	"github.com/joho/godotenv"
)

func NewEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

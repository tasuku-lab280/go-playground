package main

import (
	"github.com/tasuku-lab280/go-playground/pkg/infrastructure"
)

func main() {
	infrastructure.NewEnv()
	infrastructure.NewDB()
	infrastructure.NewRouter()
}

package main

import (
	"github.com/tasuku-lab280/go-playground/internal/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	router := infrastructure.NewRouter(db)
	router.Run()
}

package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuku-lab280/go-playground/internal/handler"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	handler.NewTodoHandler(router, db)

	return router
}

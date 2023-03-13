package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tasuku-lab280/go-playground/internal/model"
	"github.com/tasuku-lab280/go-playground/internal/repository"
	"github.com/tasuku-lab280/go-playground/internal/usecase"
	"gorm.io/gorm"
)

func NewTodoHandler(router *gin.Engine, db *gorm.DB) {
	repo := repository.NewTodoRepository(db)
	usecase := usecase.NewTodoUsecase(repo)

	// TODOの一覧取得
	router.GET("/todos", func(c *gin.Context) {
		todos, err := usecase.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)
	})

	// TODOの取得
	router.GET("/todos/:id", func(c *gin.Context) {
		var todo model.Todo
		if err := db.First(&todo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

	// TODOの作成
	router.POST("/todos", func(c *gin.Context) {
		var input model.Todo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo := model.Todo{Title: input.Title}
		db.Create(&todo)
		c.JSON(http.StatusOK, todo)
	})

	// TODOの更新
	router.PUT("/todos/:id", func(c *gin.Context) {
		var todo model.Todo
		if err := db.First(&todo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		var input model.Todo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Model(&todo).Updates(input)
		c.JSON(http.StatusOK, todo)
	})

	// TODOの削除
	router.DELETE("/todos/:id", func(c *gin.Context) {
		var todo model.Todo
		if err := db.First(&todo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		db.Delete(&todo)
		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
	})
}

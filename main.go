package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
  gorm.Model
  Title string
}

func main() {
	dsn := "user:password@tcp(go_playground_db:3306)/go_playground?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	db.AutoMigrate(&Todo{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "SUCCESS",
		})
	})

	// TODOの一覧取得
	r.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	// TODOの取得
	r.GET("/todos/:id", func(c *gin.Context) {
		var todo Todo
		if err := db.First(&todo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

	// TODOの作成
	r.POST("/todos", func(c *gin.Context) {
		var input Todo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo := Todo{Title: input.Title}
		db.Create(&todo)
		c.JSON(http.StatusOK, todo)
	})

	// TODOの更新
	r.PUT("/todos/:id", func(c *gin.Context) {
		var todo Todo
		if err := db.First(&todo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		var input Todo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Model(&todo).Updates(input)
		c.JSON(http.StatusOK, todo)
	})

	// TODOの削除
	r.DELETE("/todos/:id", func(c *gin.Context) {
		var todo Todo
		if err := db.First(&todo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		db.Delete(&todo)
		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
	})

	r.Run()
}

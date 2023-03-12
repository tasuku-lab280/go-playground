package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
	dsn := "user:password@tcp(go_playground_db:3306)/go_playground?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D42", Price: 100})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "SUCCESS",
		})
	})
	r.Run()
}

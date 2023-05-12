package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
	"log"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
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

package infrastructure

import (
	"github.com/tasuku-lab280/go-playground/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "user:password@tcp(go_playground_db:3306)/go_playground?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	db.AutoMigrate(&model.Todo{})

	return db
}

package repository

import (
	"github.com/tasuku-lab280/go-playground/internal/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll() ([]model.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) FindAll() ([]model.Todo, error) {
	var todos []model.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

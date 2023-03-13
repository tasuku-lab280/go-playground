package usecase

import (
	"github.com/tasuku-lab280/go-playground/internal/model"
	"github.com/tasuku-lab280/go-playground/internal/repository"
)

type TodoUsecase interface {
	GetAll() ([]model.Todo, error)
}

type todoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		repo: repo,
	}
}

func (u *todoUsecase) GetAll() ([]model.Todo, error) {
	return u.repo.FindAll()
}

package repository

import (
	"github.com/tasuku-lab280/go-playground/pkg/model"
	"gorm.io/gorm"
)

type IProductRepository interface {
	GetAllProducts(products *[]model.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) GetAllProducts(products *[]model.Product) error {
	if err := pr.db.Find(products).Error; err != nil {
		return err
	}

	return nil
}

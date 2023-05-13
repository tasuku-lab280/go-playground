package usecase

import (
	"github.com/tasuku-lab280/go-playground/pkg/model"
	"github.com/tasuku-lab280/go-playground/pkg/repository"
)

type IProductUsecase interface {
	GetAllProducts() ([]model.ProductResponse, error)
}

type productUsecase struct {
	pr repository.IProductRepository
}

func NewProductUsecase(pr repository.IProductRepository) IProductUsecase {
	return &productUsecase{pr}
}

func (pu *productUsecase) GetAllProducts() ([]model.ProductResponse, error) {
	products := []model.Product{}
	if err := pu.pr.GetAllProducts(&products); err != nil {
		return nil, err
	}

	res := []model.ProductResponse{}
	for _, v := range products {
		t := model.ProductResponse{
			Id:        v.Id,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		res = append(res, t)
	}

	return res, nil
}

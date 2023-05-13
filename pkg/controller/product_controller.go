package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tasuku-lab280/go-playground/pkg/usecase"
)

type IProductController interface {
	GetAllProducts(c *gin.Context)
}

type productController struct {
	pu usecase.IProductUsecase
}

func NewProductController(pu usecase.IProductUsecase) IProductController {
	return &productController{pu}
}

func (pc *productController) GetAllProducts(c *gin.Context) {
	res, err := pc.pu.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

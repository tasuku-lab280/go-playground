package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuku-lab280/go-playground/pkg/controller"
	"github.com/tasuku-lab280/go-playground/pkg/repository"
	"github.com/tasuku-lab280/go-playground/pkg/usecase"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) {
	r := gin.Default()

	pr := repository.NewProductRepository(db)
	pu := usecase.NewProductUsecase(pr)
	pc := controller.NewProductController(pu)
	r.GET("products", pc.GetAllProducts)

	r.Run()
}

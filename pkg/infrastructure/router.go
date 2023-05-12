package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "SUCCESS",
		})
	})
	r.Run()
}

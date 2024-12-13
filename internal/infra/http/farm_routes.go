package http

import (
	controller "github.com/gabmenezesdev/go-tech-challenge/internal/infra/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	farmRoutes := router.Group("api/v1/farm")
	{
		farmRoutes.POST("", controller.CreateFarmController)
		farmRoutes.GET("", controller.DeleteFarmController)
		farmRoutes.DELETE(":id", controller.GetAllFarmController)
	}
}

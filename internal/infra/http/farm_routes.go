package http

import (
	controller "github.com/gabmenezesdev/go-tech-challenge/internal/infra/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	userRoutes := router.Group("/farm")
	{
		userRoutes.POST("", controller.CreateFarmController)
		userRoutes.GET("", controller.DeleteFarmController)
		userRoutes.DELETE(":id", controller.GetAllFarmController)
	}
}

package http

import (
	controller "github.com/gabmenezesdev/go-tech-challenge/internal/infra/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	createFarmController := controller.NewCreateFarmController()
	deleteFarmController := controller.NewDeleteFarmController()
	getAllFarmController := controller.NewGetAllFarmController()

	farmRoutes := router.Group("api/v1/farm")
	{
		farmRoutes.GET("", getAllFarmController.Handle)
		farmRoutes.POST("", createFarmController.Handle)
		farmRoutes.DELETE(":id", deleteFarmController.Handle)
	}
}

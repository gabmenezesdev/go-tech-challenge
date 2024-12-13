package controller

import (
	"net/http"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/application/use-case"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
	"github.com/gin-gonic/gin"
)

func DeleteFarmController(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID is required",
		})
		return
	}

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}

	deleteFarmUseCase, err := farm.NewDeleteFarmUseCase(farmMongoDbAdapter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to initialize farm deletion use case",
		})
		return
	}

	err = deleteFarmUseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occurred during farm deletion",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Farm deleted successfully!",
	})
}

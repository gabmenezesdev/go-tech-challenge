package controller

import (
	"net/http"

	usecase "github.com/gabmenezesdev/go-tech-challenge/internal/application/use-case"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
	"github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DeleteFarmController(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		shared.LoggerError("Farm ID is required", nil)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID is required",
			"details": "id not informed",
		})
		return
	}

	shared.LoggerInfo("Received request to delete farm", zap.String("farmId", id))

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}

	deleteFarmUseCase, err := usecase.NewDeleteFarmUseCase(farmMongoDbAdapter)
	if err != nil {
		shared.LoggerError("Unable to initialize farm deletion use case", err, zap.String("farmId", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to initialize farm deletion use case",
			"details": err.Error(),
		})
		return
	}

	err = deleteFarmUseCase.Execute(id)
	if err != nil {
		shared.LoggerError("Error occurred during farm deletion", err, zap.String("farmId", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occurred during farm deletion",
			"details": err.Error(),
		})
		return
	}

	shared.LoggerInfo("Farm deleted successfully", zap.String("farmId", id))
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Farm deleted successfully!",
	})
}

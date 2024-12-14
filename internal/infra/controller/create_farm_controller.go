package controller

import (
	"net/http"

	usecase "github.com/gabmenezesdev/go-tech-challenge/internal/application/use-case"
	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
	"github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateFarmController(ctx *gin.Context) {
	shared.LoggerInfo("Received request to create farm")

	var requestBody farm.FarmDto
	if err := ctx.BindJSON(&requestBody); err != nil {
		shared.LoggerError("Invalid request payload", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	shared.LoggerInfo("Request payload parsed successfully", zap.String("farmName", requestBody.Name))

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}
	cropMongoDbAdapter := repository.CropRepositoryMongoAdapter{}

	createFarmUseCase, err := usecase.NewCreateFarmUseCase(farmMongoDbAdapter, cropMongoDbAdapter)
	if err != nil {
		shared.LoggerError("Error initializing CreateFarmUseCase", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error initializing farm creation process",
		})
		return
	}

	err = createFarmUseCase.Execute(
		requestBody.Name,
		requestBody.LandArea,
		requestBody.Unit,
		requestBody.Address,
		requestBody.Crops,
	)
	if err != nil {
		shared.LoggerError("Error executing CreateFarmUseCase", err, zap.String("farmName", requestBody.Name))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating farm",
			"details": err.Error(),
		})
		return
	}

	shared.LoggerInfo("Farm created successfully", zap.String("farmName", requestBody.Name))
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Farm created successfully!",
	})
}

package controller

import (
	"net/http"

	usecase "github.com/gabmenezesdev/go-tech-challenge/internal/application/use-case"
	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
	"github.com/gin-gonic/gin"
)

func CreateFarmController(ctx *gin.Context) {
	var requestBody farm.FarmDto

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}
	cropMongoDbAdapter := repository.CropRepositoryMongoAdapter{}

	createFarmUseCase, err := usecase.NewCreateFarmUseCase(farmMongoDbAdapter, cropMongoDbAdapter)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Error during Farm Creation",
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating farm",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Farm created successfully!",
	})

}

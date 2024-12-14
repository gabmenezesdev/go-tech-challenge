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

func NewCreateFarmController() *createFarmController {
	return &createFarmController{}
}

type createFarmController struct{}

// CreateFarm Creates a new farm
// @Summary Create a new farm
// @Description This endpoint allows the creation of a new farm by providing the necessary details.
// @Tags Farm
// @Accept json
// @Produce json
// @Param requestBody body farm.CreateFarmDto true "Farm Information"
// @Success 201 {object} shared.SuccessResponse "Successfully created farm"
// @Failure 400 {object} shared.ErrorResponse "Bad Request: Invalid request payload"
// @Failure 500 {object} shared.ErrorResponse "Internal Server Error: Error initializing or executing farm creation process"
// @Router /farm [post]
func (cfc *createFarmController) Handle(ctx *gin.Context) {
	shared.LoggerInfo("Received request to create farm")

	var requestBody farm.CreateFarmDto
	if err := ctx.BindJSON(&requestBody); err != nil {
		shared.LoggerError("Invalid request payload", err)
		ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Message: "Invalid request payload"})
		return
	}

	shared.LoggerInfo("Request payload parsed successfully", zap.String("farmName", requestBody.Name))

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}
	cropMongoDbAdapter := repository.CropRepositoryMongoAdapter{}

	createFarmUseCase, err := usecase.NewCreateFarmUseCase(farmMongoDbAdapter, cropMongoDbAdapter)
	if err != nil {
		shared.LoggerError("Error initializing CreateFarmUseCase", err)
		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{Message: "Error initializing farm creation process"})
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
		ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Message: "Error creating farm", Details: err.Error()})
		return
	}

	shared.LoggerInfo("Farm created successfully", zap.String("farmName", requestBody.Name))
	ctx.JSON(http.StatusCreated, shared.SuccessResponse{
		Message: "Farm created successfully!",
	})
}

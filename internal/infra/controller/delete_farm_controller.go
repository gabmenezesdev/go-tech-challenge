package controller

import (
	"net/http"

	usecase "github.com/gabmenezesdev/go-tech-challenge/internal/application/use-case"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
	"github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewDeleteFarmController() *deleteFarmController {
	return &deleteFarmController{}
}

type deleteFarmController struct{}

// DeleteFarm Deletes a farm by its ID
// @Summary Delete a farm by its ID
// @Description This endpoint allows the deletion of an existing farm by providing its unique ID. If the farm is found, it will be deleted. If the farm is not found or an error occurs during deletion, appropriate error responses are returned.
// @Tags Farm
// @Accept json
// @Produce json
// @Param id path string true "Farm ID"
// @Success 200 {object} shared.SuccessResponse "Successfully deleted farm"
// @Failure 400 {object} shared.ErrorResponse "Bad Request: Farm ID is required or invalid"
// @Failure 500 {object} shared.ErrorResponse "Internal Server Error: Error initializing or executing farm deletion"
// @Router /farm/{id} [delete]
func (dfc *deleteFarmController) Handle(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		shared.LoggerError("Farm ID is required", nil)
		ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{
			Message: "ID is required",
			Details: "id not informed",
		})
		return
	}

	shared.LoggerInfo("Received request to delete farm", zap.String("farmId", id))

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}

	deleteFarmUseCase, err := usecase.NewDeleteFarmUseCase(farmMongoDbAdapter)
	if err != nil {
		shared.LoggerError("Unable to initialize farm deletion use case", err, zap.String("farmId", id))
		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{
			Message: "Unable to initialize farm deletion use case",
		})
		return
	}

	err = deleteFarmUseCase.Execute(id)
	if err != nil {
		shared.LoggerError("Error occurred during farm deletion", err, zap.String("farmId", id))
		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{
			Message: "Error occurred during farm deletion",
			Details: err.Error(),
		})
		return
	}

	shared.LoggerInfo("Farm deleted successfully", zap.String("farmId", id))
	ctx.JSON(http.StatusOK, shared.SuccessResponse{
		Message: "Farm deleted successfully!",
	})
}

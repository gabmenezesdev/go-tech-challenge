package controller

import (
	"net/http"
	"strconv"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/application/use-case"
	dao "github.com/gabmenezesdev/go-tech-challenge/internal/infra/DAO"
	"github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewGetAllFarmController() *getAllFarmController {
	return &getAllFarmController{}
}

type getAllFarmController struct{}

// GetAllFarms Retrieves all farms with optional filters
// @Summary Retrieve a list of farms with optional filters
// @Description This endpoint allows the retrieval of a list of farms, with optional query parameters for filtering by name, unit, crop type, and land area. The response supports pagination with the 'skip' and 'perpage' parameters. If no farms are found, an empty array will be returned. If the required parameters are missing or invalid, appropriate error responses will be returned.
// @Tags Farm
// @Accept json
// @Produce json
// @Param skip query int false "Number of items to skip for pagination"
// @Param perpage query int false "Number of items per page for pagination"
// @Param name query string false "Name of the farm"
// @Param unit query string false "Unit for land area"
// @Param croptype query string false "Type of crop grown on the farm"
// @Param landareainit query string false "Initial land area for filtering"
// @Param landarealimit query string false "Limit land area for filtering"
// @Success 200 {object} shared.SuccessResponse "List of farms retrieved successfully"
// @Failure 400 {object} shared.ErrorResponse "Bad Request: Missing or invalid query parameters"
// @Failure 500 {object} shared.ErrorResponse "Internal Server Error: Error retrieving farms"
// @Router /farm [get]
func (gaf *getAllFarmController) Handle(ctx *gin.Context) {
	skipStr := ctx.DefaultQuery("skip", "")
	perPageStr := ctx.DefaultQuery("perpage", "")
	name := ctx.DefaultQuery("name", "")
	unit := ctx.DefaultQuery("unit", "")
	croptype := ctx.DefaultQuery("croptype", "")
	landAreaInit := ctx.DefaultQuery("landareainit", "")
	landAreaLimit := ctx.DefaultQuery("landarealimit", "")

	if skipStr == "" || perPageStr == "" {
		shared.LoggerError("Missing query parameters", nil, zap.String("skip", skipStr), zap.String("perpage", perPageStr))
		ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{
			Message: "Both 'skip' and 'perpage' query parameters are required",
			Details: "'skip' or 'perpage' not informed",
		})
		return
	}

	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		shared.LoggerError("Error converting 'skip' parameter", err, zap.String("skip", skipStr))
		ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{
			Message: "'skip' must be an integer",
			Details: err.Error(),
		})
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		shared.LoggerError("Error converting 'perpage' parameter", err, zap.String("perpage", perPageStr))
		ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{
			Message: "'perpage' must be an integer",
			Details: err.Error(),
		})
		return
	}

	farmMongoDbAdapterDAO := dao.FarmDaoMongoDB{}
	NewGetAllFarmsUseCase, err := farm.NewGetAllFarmsUseCase(farmMongoDbAdapterDAO)
	if err != nil {
		shared.LoggerError("Error initializing farm retrieval use case", err)
		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{
			Message: "Error during farm retrieval",
			Details: err.Error(),
		})
		return
	}

	farmFilters := dao.FarmFilters{
		Name:          name,
		Unit:          unit,
		CropType:      croptype,
		LandAreaLimit: landAreaLimit,
		LandAreaInit:  landAreaInit,
	}

	foundFarms, err := NewGetAllFarmsUseCase.Execute(skip, perPage, farmFilters)
	if err != nil {
		shared.LoggerError("Error retrieving farms", err)
		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{
			Message: "Error during farm get",
			Details: err.Error(),
		})
		return
	}

	shared.LoggerInfo("Farms retrieved successfully", zap.Int("totalFarms", len(foundFarms)))

	if len(foundFarms) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No farms found!",
			"data":    [0]int{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Farms retrieved successfully!",
		"data":    foundFarms,
	})
}

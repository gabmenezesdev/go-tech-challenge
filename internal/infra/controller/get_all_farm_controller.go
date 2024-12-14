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

func (gaf *getAllFarmController) Handle(ctx *gin.Context) {
	skipStr := ctx.DefaultQuery("skip", "")
	perPageStr := ctx.DefaultQuery("perpage", "")
	name := ctx.DefaultQuery("name", "")
	unit := ctx.DefaultQuery("skip", "")
	croptype := ctx.DefaultQuery("croptype", "")
	isirrigated := ctx.DefaultQuery("isirrigated", "")
	isinsured := ctx.DefaultQuery("isinsured", "")

	if skipStr == "" || perPageStr == "" {
		shared.LoggerError("Missing query parameters", nil, zap.String("skip", skipStr), zap.String("perpage", perPageStr))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Both 'skip' and 'perpage' query parameters are required",
			"details": "'skip' or 'perpage' not informed",
		})
		return
	}

	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		shared.LoggerError("Error converting 'skip' parameter", err, zap.String("skip", skipStr))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "'skip' must be an integer",
			"details": err.Error(),
		})
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		shared.LoggerError("Error converting 'perpage' parameter", err, zap.String("perpage", perPageStr))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "'perpage' must be an integer",
			"details": err.Error(),
		})
		return
	}

	farmMongoDbAdapterDAO := dao.FarmDaoMongoDB{}
	NewGetAllFarmsUseCase, err := farm.NewGetAllFarmsUseCase(farmMongoDbAdapterDAO)
	if err != nil {
		shared.LoggerError("Error initializing farm retrieval use case", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error during farm retrieval",
			"details": err.Error(),
		})
		return
	}

	farmFilters := dao.FarmFilters{
		Name:        name,
		Unit:        unit,
		CropType:    croptype,
		IsIrrigated: isirrigated,
		IsInsured:   isinsured,
	}

	foundFarms, err := NewGetAllFarmsUseCase.Execute(skip, perPage, farmFilters)
	if err != nil {
		shared.LoggerError("Error retrieving farms", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error during farm get",
			"details": err.Error(),
		})
		return
	}

	shared.LoggerInfo("Farms retrieved successfully", zap.Int("totalFarms", len(foundFarms)))

	ctx.JSON(http.StatusOK, gin.H{
		"skip":    skip,
		"perpage": perPage,
		"message": "Farms got successfully!",
		"data":    foundFarms,
	})
}

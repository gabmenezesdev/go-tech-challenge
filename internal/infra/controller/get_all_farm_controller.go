package controller

import (
	"net/http"
	"strconv"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/application/use-case"
	dao "github.com/gabmenezesdev/go-tech-challenge/internal/infra/DAO"
	"github.com/gin-gonic/gin"
)

func GetAllFarmController(ctx *gin.Context) {
	skipStr := ctx.DefaultQuery("skip", "")
	perPageStr := ctx.DefaultQuery("perpage", "")
	name := ctx.DefaultQuery("name", "")
	unit := ctx.DefaultQuery("skip", "")
	croptype := ctx.DefaultQuery("croptype", "")
	isirrigated := ctx.DefaultQuery("isirrigated", "")
	isinsured := ctx.DefaultQuery("isinsured", "")

	if skipStr == "" || perPageStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Both 'skip' and 'perpage' query parameters are required",
		})
		return
	}

	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "'skip' must be an integer",
		})
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "'perpage' must be an integer",
		})
		return
	}

	farmMongoDbAdapterDAO := dao.FarmDaoMongoDB{}
	NewGetAllFarmsUseCase, err := farm.NewGetAllFarmsUseCase(farmMongoDbAdapterDAO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error during farm retrieval",
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error during farm get",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"skip":    skip,
		"perpage": perPage,
		"message": "Farms got successfully!",
		"data":    foundFarms,
	})
}

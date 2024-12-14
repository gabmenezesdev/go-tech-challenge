package usecase

import (
	"github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	dao "github.com/gabmenezesdev/go-tech-challenge/internal/infra/DAO"
	"github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"go.uber.org/zap"
)

type GetAllFarm struct {
	farmDao dao.FarmDao
}

func NewGetAllFarmsUseCase(farmDao dao.FarmDao) (*GetAllFarm, error) {
	shared.LoggerInfo("Initializing GetAllFarms use case")
	return &GetAllFarm{
		farmDao: farmDao,
	}, nil
}

func (gaf *GetAllFarm) Execute(skip int, perPage int, filters dao.FarmFilters) ([]farm.FarmDto, error) {
	shared.LoggerInfo("Executing GetAllFarms use case",
		zap.Int("skip", skip),
		zap.Int("perPage", perPage),
		zap.Any("filters", filters),
	)

	allFarms, err := gaf.farmDao.GetAllFarms(skip, perPage, filters)
	if err != nil {
		shared.LoggerError("Error retrieving farms", err, zap.Int("skip", skip), zap.Int("perPage", perPage), zap.Any("filters", filters))
		return []farm.FarmDto{}, err
	}

	shared.LoggerInfo("Farms retrieved successfully", zap.Int("count", len(allFarms)))
	return allFarms, nil
}

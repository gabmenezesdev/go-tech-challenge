package usecase

import (
	"github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	dao "github.com/gabmenezesdev/go-tech-challenge/internal/infra/DAO"
)

type GetAllFarm struct {
	farmDao dao.FarmDao
}

func NewGetAllFarmsUseCase(farmDao dao.FarmDao) (*GetAllFarm, error) {
	return &GetAllFarm{
		farmDao: farmDao,
	}, nil
}

func (gaf *GetAllFarm) Execute(skip int, perPage int) ([]farm.FarmDto, error) {
	allFarms, err := gaf.farmDao.GetAllFarms(skip, perPage)
	if err != nil {
		return []farm.FarmDto{}, err
	}

	return allFarms, nil
}

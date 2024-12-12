package usecase

import (
	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
)

type CreateFarm struct {
	farmRepository repository.FarmRepository
}

func NewCreateFarmUseCase(farmRepository repository.FarmRepository) (*CreateFarm, error) {
	return &CreateFarm{
		farmRepository: farmRepository,
	}, nil
}

func (cf *CreateFarm) Execute(name string, landArea float64, unit string, address string, crops []string) error {

	newFarm, err := farm.NewFarm(name, landArea, unit, address, crops)
	if err != nil {
		return err
	}

	err = cf.farmRepository.CreateFarm(newFarm)
	if err != nil {
		return err
	}
	return nil
}

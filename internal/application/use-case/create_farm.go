package usecase

import (
	crop "github.com/gabmenezesdev/go-tech-challenge/internal/domain/crop"
	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
)

type CreateFarm struct {
	farmRepository repository.FarmRepository
	cropRepository repository.CropRepository
}

func NewCreateFarmUseCase(farmRepository repository.FarmRepository, cropRepository repository.CropRepository) (*CreateFarm, error) {
	return &CreateFarm{
		farmRepository: farmRepository,
		cropRepository: cropRepository,
	}, nil
}

func (cf *CreateFarm) Execute(name string, landArea float64, unit string, address string, crops []crop.CropDto) error {

	newFarm, err := farm.NewFarm(name, landArea, unit, address)
	if err != nil {
		return err
	}
	createdFarmId, err := cf.farmRepository.CreateFarm(newFarm)
	if err != nil {
		return err
	}

	for _, production := range crops {
		newCrop, err := crop.NewCrop(production.CropType, production.IsIrrigated, production.IsInsured)
		cf.cropRepository.CreateCrop(newCrop, createdFarmId)
		if err != nil {
			return err
		}
	}

	return nil
}

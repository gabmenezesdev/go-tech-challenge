package usecase

import (
	crop "github.com/gabmenezesdev/go-tech-challenge/internal/domain/crop"
	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
	"github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"go.uber.org/zap"
)

type createFarm struct {
	farmRepository repository.FarmRepository
	cropRepository repository.CropRepository
}

func NewCreateFarmUseCase(farmRepository repository.FarmRepository, cropRepository repository.CropRepository) (*createFarm, error) {
	shared.LoggerInfo("Initializing CreateFarm use case")
	return &createFarm{
		farmRepository: farmRepository,
		cropRepository: cropRepository,
	}, nil
}

func (cf *createFarm) Execute(name string, landArea float64, unit string, address string, crops []crop.CropDto) error {
	shared.LoggerInfo("Starting create farm process", zap.String("farmName", name))

	newFarm, err := farm.NewFarm(name, landArea, unit, address)
	if err != nil {
		shared.LoggerError("Failed to create farm instance", err)
		return err
	}

	createdFarmId, err := cf.farmRepository.CreateFarm(newFarm)
	if err != nil {
		shared.LoggerError("Error saving farm to repository", err, zap.String("farmName", name))
		return err
	}

	shared.LoggerInfo("Farm created successfully", zap.String("farmId", createdFarmId))

	for _, production := range crops {
		newCrop, err := crop.NewCrop(production.CropType, production.IsIrrigated, production.IsInsured)
		if err != nil {
			shared.LoggerError("Failed to create crop instance", err, zap.String("cropType", production.CropType))
			return err
		}

		err = cf.cropRepository.CreateCrop(newCrop, createdFarmId)
		if err != nil {
			shared.LoggerError("Error saving crop to repository", err, zap.String("cropType", production.CropType), zap.String("farmId", createdFarmId))
			return err
		}

		shared.LoggerInfo("Crop added successfully", zap.String("cropType", production.CropType), zap.String("farmId", createdFarmId))
	}

	shared.LoggerInfo("Farm creation process completed successfully", zap.String("farmId", createdFarmId))
	return nil
}

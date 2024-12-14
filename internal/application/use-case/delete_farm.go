package usecase

import (
	"fmt"

	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
	"github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type deleteFarm struct {
	farmRepository repository.FarmRepository
}

func NewDeleteFarmUseCase(farmRepository repository.FarmRepository) (*deleteFarm, error) {
	shared.LoggerInfo("Initializing DeleteFarm use case")
	return &deleteFarm{
		farmRepository: farmRepository,
	}, nil
}

func (df *deleteFarm) Execute(farmId string) error {
	shared.LoggerInfo("Executing DeleteFarm use case", zap.String("farmId", farmId))

	_, err := df.farmRepository.GetFarmById(farmId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			shared.LoggerError("Farm not found", err, zap.String("farmId", farmId))
			return fmt.Errorf("farm not found with id: %s", farmId)
		}
		shared.LoggerError("Error while retrieving farm", err, zap.String("farmId", farmId))
		return err
	}

	shared.LoggerInfo("Farm found, proceeding with deletion", zap.String("farmId", farmId))

	err = df.farmRepository.DeleteFarmById(farmId)
	if err != nil {
		shared.LoggerError("Error while deleting farm", err, zap.String("farmId", farmId))
		return err
	}

	shared.LoggerInfo("Farm deleted successfully", zap.String("farmId", farmId))
	return nil
}

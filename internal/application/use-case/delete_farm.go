package usecase

import (
	"fmt"

	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteFarm struct {
	farmRepository repository.FarmRepository
}

func NewDeleteFarmUseCase(farmRepository repository.FarmRepository) (*DeleteFarm, error) {
	return &DeleteFarm{
		farmRepository: farmRepository,
	}, nil
}

func (df *DeleteFarm) Execute(farmId string) error {
	_, err := df.farmRepository.GetFarmById(farmId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("farm not found with id: %s", farmId)
		}
		return err
	}

	err = df.farmRepository.DeleteFarmById(farmId)
	if err != nil {
		return err
	}

	return nil
}

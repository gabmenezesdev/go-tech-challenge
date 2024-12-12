package usecase

import (
	"testing"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
)

func TestDeleteFarm(t *testing.T) {
	t.Setenv("MONGODB_URL", "mongodb://localhost:27017")
	t.Setenv("MONGODB_DATABASE", "teste")

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}

	DeleteFarmUseCase, err := NewDeleteFarmUseCase(farmMongoDbAdapter)

	newFarm, err := farm.NewFarm("MyFarm", 100.5, "acre", "123 Farm Address")
	if err != nil {
		t.Error("Need to create a farm before delete")
	}

	createdFarmId, err := farmMongoDbAdapter.CreateFarm(newFarm)
	if err != nil {
		t.Error("Need to create a farm before delete, but got error")
	}

	err = DeleteFarmUseCase.Execute(createdFarmId)
	if err != nil {
		t.Error("Error on deleting farm")
	}
}

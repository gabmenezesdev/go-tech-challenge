package usecase

import (
	"fmt"
	"testing"

	crop "github.com/gabmenezesdev/go-tech-challenge/internal/domain/crop"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
)

func TestCreateFarm(t *testing.T) {
	t.Setenv("MONGODB_URL", "mongodb://localhost:27017")
	t.Setenv("MONGODB_DATABASE", "teste")

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}
	cropMongoDbAdapter := repository.CropRepositoryMongoAdapter{}

	createFarmUseCase, err := NewCreateFarmUseCase(farmMongoDbAdapter, cropMongoDbAdapter)
	if err != nil {
		t.Error("Should Create a farm in database, but got error")
	}

	crops := []crop.CropDto{
		{
			CropType:    "RICE",
			IsIrrigated: true,
			IsInsured:   true,
		},
		{
			CropType:    "BEANS",
			IsIrrigated: false,
			IsInsured:   true,
		},
	}
	err = createFarmUseCase.Execute("MyFarm", 100.5, "acre", "123 Farm Address", crops)
	if err != nil {
		fmt.Println(err)
		t.Error("Should Create a farm in database, but got error")
	}
}

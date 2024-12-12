package usecase

import (
	"fmt"
	"testing"

	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
)

func TestCreateFarm(t *testing.T) {
	t.Setenv("MONGODB_URL", "mongodb://localhost:27017")
	t.Setenv("MONGODB_DATABASE", "teste")

	mongoDbAdapter := repository.FarmRepositoryMongoAdapter{}
	createFarmUseCase, err := NewCreateFarmUseCase(mongoDbAdapter)
	if err != nil {
		t.Error("Should Create a farm in database, but got error")
	}

	crops := []string{"corn", "wheat"}
	err = createFarmUseCase.Execute("MyFarm", 100.5, "acre", "123 Farm Address", crops)
	if err != nil {
		fmt.Println(err)
		t.Error("Should Create a farm in database, but got error")
	}
}

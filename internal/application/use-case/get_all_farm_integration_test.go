package usecase

import (
	"testing"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	dao "github.com/gabmenezesdev/go-tech-challenge/internal/infra/DAO"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
)

func TestGetAllFarm(t *testing.T) {
	t.Setenv("MONGODB_URL", "mongodb://localhost:27017")
	t.Setenv("MONGODB_DATABASE", "teste")

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}
	farmMongoDbAdapterDAO := dao.FarmDaoMongoDB{}

	NewGetAllFarmsUseCase, err := NewGetAllFarmsUseCase(farmMongoDbAdapterDAO)
	if err != nil {
		t.Error(err)
	}

	firstFarm, err := farm.NewFarm("FirstFarm", 10.5, "hectare", "123 Farm Address")
	if err != nil {
		t.Error(err)
	}

	secondFarm, err := farm.NewFarm("SecondFarm", 20.7, "acre", "20 Address new")
	if err != nil {
		t.Error(err)
	}

	thirdFarm, err := farm.NewFarm("ThirdFarm", 30.8, "square_km", "15 Faria Lima")
	if err != nil {
		t.Error(err)
	}

	_, err = farmMongoDbAdapter.CreateFarm(firstFarm)
	if err != nil {
		t.Error(err)
	}

	_, err = farmMongoDbAdapter.CreateFarm(secondFarm)
	if err != nil {
		t.Error(err)
	}

	_, err = farmMongoDbAdapter.CreateFarm(thirdFarm)
	if err != nil {
		t.Error(err)
	}
	farmFilter := dao.FarmFilters{}
	foundFarms, err := NewGetAllFarmsUseCase.Execute(0, 2, farmFilter)
	if err != nil {
		t.Error(err)
	}

	if len(foundFarms) == 0 {
		t.Error("Expected to have an array of Farms, but got none")
	}
}

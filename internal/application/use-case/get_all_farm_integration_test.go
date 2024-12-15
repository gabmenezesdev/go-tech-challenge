package usecase

import (
	"testing"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	dao "github.com/gabmenezesdev/go-tech-challenge/internal/infra/DAO"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/repository"
)

func TestGetAllFarmWithFilters(t *testing.T) {
	t.Setenv("MONGODB_URL", "mongodb://localhost:27017")
	t.Setenv("MONGODB_DATABASE", "teste")

	farmMongoDbAdapter := repository.FarmRepositoryMongoAdapter{}
	farmMongoDbAdapterDAO := dao.FarmDaoMongoDB{}

	NewGetAllFarmsUseCase, err := NewGetAllFarmsUseCase(farmMongoDbAdapterDAO)
	if err != nil {
		t.Error(err)
	}

	firstFarm, err := farm.NewFarm("AppleFarm", 15.5, "hectare", "10 Orchard Lane")
	if err != nil {
		t.Error(err)
	}

	secondFarm, err := farm.NewFarm("BerryFarm", 25.7, "acre", "25 Berry Street")
	if err != nil {
		t.Error(err)
	}

	thirdFarm, err := farm.NewFarm("CherryFarm", 35.8, "square_km", "35 Cherry Drive")
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

	tests := []struct {
		name      string
		filters   dao.FarmFilters
		expectLen int
	}{
		{
			name: "Filter by Name",
			filters: dao.FarmFilters{
				Name: "AppleFarm",
			},
			expectLen: 1,
		},
		{
			name: "Filter by Unit",
			filters: dao.FarmFilters{
				Unit: "acre",
			},
			expectLen: 1,
		},
		{
			name: "Filter by Land Area Less Than",
			filters: dao.FarmFilters{
				LandAreaLimit: "30",
			},
			expectLen: 2,
		},
		{
			name: "Filter by Land Area Greater Than",
			filters: dao.FarmFilters{
				LandAreaInit: "20",
			},
			expectLen: 2,
		},
		{
			name: "Filter by Land Area Range",
			filters: dao.FarmFilters{
				LandAreaInit:  "20",
				LandAreaLimit: "40",
			},
			expectLen: 2,
		},
		{
			name:      "No Filters",
			filters:   dao.FarmFilters{},
			expectLen: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foundFarms, err := NewGetAllFarmsUseCase.Execute(0, 10, tt.filters)
			if err != nil {
				t.Error(err)
			}

			if len(foundFarms) != tt.expectLen {
				t.Errorf("Expected %d farms, but got %d", tt.expectLen, len(foundFarms))
			}
		})
	}
}

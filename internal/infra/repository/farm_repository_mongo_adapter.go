package repository

import (
	"context"
	"fmt"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/database"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	FARM_SCHEMA = "farms"
)

type FarmRepositoryMongoAdapter struct{}

func (f FarmRepositoryMongoAdapter) CreateFarm(u *farm.Farm) error {
	fmt.Println("entrou aqui")
	client, err := database.InitConnection()
	if err != nil {
		return err
	}

	_, err = client.Collection(FARM_SCHEMA).InsertOne(context.Background(), bson.M{
		"name":      u.GetName(),
		"land_area": u.GetLandArea(),
		"unit":      u.GetUnit(),
		"address":   u.GetAddress(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (f FarmRepositoryMongoAdapter) DeleteFarmById(id int64) error {
	return nil
}

func (f FarmRepositoryMongoAdapter) GetFarmById(id int64) error {
	return nil
}

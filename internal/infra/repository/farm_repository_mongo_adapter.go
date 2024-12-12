package repository

import (
	"context"
	"fmt"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FarmRepositoryMongoAdapter struct{}

func (f FarmRepositoryMongoAdapter) CreateFarm(farm *farm.Farm) (string, error) {
	fmt.Println("entrou aqui")
	client, err := database.InitConnection()
	if err != nil {
		return "", err
	}

	res, err := client.Collection(FARM_SCHEMA).InsertOne(context.Background(), bson.M{
		"name":      farm.GetName(),
		"land_area": farm.GetLandArea(),
		"unit":      farm.GetUnit(),
		"address":   farm.GetAddress(),
		"crops":     []interface{}{},
	})
	if err != nil {
		return "", err
	}

	id := res.InsertedID

	return id.(primitive.ObjectID).Hex(), nil
}

func (f FarmRepositoryMongoAdapter) DeleteFarmById(id int64) error {
	return nil
}

func (f FarmRepositoryMongoAdapter) GetFarmById(id int64) error {
	return nil
}

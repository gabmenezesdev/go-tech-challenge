package repository

import (
	"context"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FarmRepositoryMongoAdapter struct{}

func (f FarmRepositoryMongoAdapter) CreateFarm(farm *farm.Farm) (string, error) {
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

func (f FarmRepositoryMongoAdapter) DeleteFarmById(farmId string) error {
	client, err := database.InitConnection()
	if err != nil {
		return err
	}
	objectID, err := primitive.ObjectIDFromHex(farmId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = client.Collection(FARM_SCHEMA).DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (f FarmRepositoryMongoAdapter) GetFarmById(farmId string) (farm.Farm, error) {
	client, err := database.InitConnection()
	if err != nil {
		return farm.Farm{}, err
	}
	objectID, err := primitive.ObjectIDFromHex(farmId)
	if err != nil {
		return farm.Farm{}, err
	}

	filter := bson.M{"_id": objectID}

	var farmData farm.Farm
	err = client.Collection(FARM_SCHEMA).FindOne(context.Background(), filter).Decode(&farmData)
	if err != nil {
		if err != nil {
			return farm.Farm{}, err
		}
	}

	return farmData, nil
}

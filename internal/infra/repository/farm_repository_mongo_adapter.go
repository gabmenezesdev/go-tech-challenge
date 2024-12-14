package repository

import (
	"context"
	"log"

	farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/database"
	shared "github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FarmRepositoryMongoAdapter struct{}

func (f FarmRepositoryMongoAdapter) CreateFarm(farm *farm.Farm) (string, error) {
	client, err := database.InitConnection()
	if err != nil {
		shared.LoggerError("Error initializing database connection", err)
		return "", err
	}

	shared.LoggerInfo("Database connection established")

	res, err := client.Collection(shared.FARM_SCHEMA).InsertOne(context.Background(), bson.M{
		"name":      farm.GetName(),
		"land_area": farm.GetLandArea(),
		"unit":      farm.GetUnit(),
		"address":   farm.GetAddress(),
		"crops":     []interface{}{},
	})
	if err != nil {
		shared.LoggerError("Error inserting farm into database", err)
		return "", err
	}

	id := res.InsertedID

	if err := database.CloseConnection(); err != nil {
		shared.LoggerError("Failed to close database connection", err)
		log.Fatalf("Failed to close database connection: %v", err)
	}

	shared.LoggerInfo("Database connection closed")

	return id.(primitive.ObjectID).Hex(), nil
}

func (f FarmRepositoryMongoAdapter) DeleteFarmById(farmId string) error {
	client, err := database.InitConnection()
	if err != nil {
		shared.LoggerError("Error initializing database connection", err)
		return err
	}

	shared.LoggerInfo("Database connection established")

	objectID, err := primitive.ObjectIDFromHex(farmId)
	if err != nil {
		shared.LoggerError("Error converting farmId to ObjectID", err)
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = client.Collection(shared.FARM_SCHEMA).DeleteOne(context.Background(), filter)
	if err != nil {
		shared.LoggerError("Error deleting farm from database", err)
		return err
	}

	shared.LoggerInfo("Farm deleted successfully")

	if err := database.CloseConnection(); err != nil {
		shared.LoggerError("Failed to close database connection", err)
		log.Fatalf("Failed to close database connection: %v", err)
	}

	shared.LoggerInfo("Database connection closed")

	return nil
}

func (f FarmRepositoryMongoAdapter) GetFarmById(farmId string) (farm.Farm, error) {
	client, err := database.InitConnection()
	if err != nil {
		shared.LoggerError("Error initializing database connection", err)
		return farm.Farm{}, err
	}

	shared.LoggerInfo("Database connection established")

	objectID, err := primitive.ObjectIDFromHex(farmId)
	if err != nil {
		shared.LoggerError("Error converting farmId to ObjectID", err)
		return farm.Farm{}, err
	}

	filter := bson.M{"_id": objectID}

	var farmData farm.Farm
	err = client.Collection(shared.FARM_SCHEMA).FindOne(context.Background(), filter).Decode(&farmData)
	if err != nil {
		shared.LoggerError("Error finding farm by ID in database", err)
		return farm.Farm{}, err
	}

	shared.LoggerInfo("Farm retrieved successfully")

	if err := database.CloseConnection(); err != nil {
		shared.LoggerError("Failed to close database connection", err)
		log.Fatalf("Failed to close database connection: %v", err)
	}

	shared.LoggerInfo("Database connection closed")

	return farmData, nil
}

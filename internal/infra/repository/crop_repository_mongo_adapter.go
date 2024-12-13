package repository

import (
	"context"
	"log"

	crop "github.com/gabmenezesdev/go-tech-challenge/internal/domain/crop"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/database"
	shared "github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CropRepositoryMongoAdapter struct{}

func (c CropRepositoryMongoAdapter) CreateCrop(crop *crop.Crop, farmId string) error {
	client, err := database.InitConnection()
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(farmId)
	if err != nil {
		return err
	}

	// I decided to use addToSet because, unlike SQL, MongoDB recommends storing it nested within a single schema.
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$push": bson.M{
			"crops": bson.M{
				"crop_type":    crop.GetCropType(),
				"is_irrigated": crop.GetIsIrrigated(),
				"is_insured":   crop.GetIsInsured(),
			},
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err = client.Collection(shared.FARM_SCHEMA).UpdateOne(context.Background(), filter, update, opts)

	if err != nil {
		return err
	}

	if err := database.CloseConnection(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}

	return nil
}

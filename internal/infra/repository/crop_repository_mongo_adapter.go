package repository

import (
	"context"
	"fmt"

	crop "github.com/gabmenezesdev/go-tech-challenge/internal/domain/crop"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	FARM_SCHEMA = "farms"
)

type CropRepositoryMongoAdapter struct{}

func (c CropRepositoryMongoAdapter) CreateCrop(crop *crop.Crop, farmId string) error {
	client, err := database.InitConnection()
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(farmId)

	fmt.Println("crop")
	fmt.Println(crop)

	// I decided to do a addToSet because different from sql mongoDb recomends to store it nested in one schema
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

	_, err = client.Collection(FARM_SCHEMA).UpdateOne(context.Background(), filter, update, opts)

	if err != nil {
		return err
	}

	return nil
}

func (c CropRepositoryMongoAdapter) DeleteCropById(id int64) error {
	return nil
}

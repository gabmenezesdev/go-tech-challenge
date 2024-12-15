package dao

import (
	"context"
	"regexp"
	"strconv"

	"github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/database"
	shared "github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// Using the CQRS pattern, we create a separate query as it is more optimized for finding data.
// Since the query doesn't manipulate a domain (i.e., it doesn't receive, process, or modify entity), a DAO is more appropriate.
type FarmDaoMongoDB struct{}

func (f FarmDaoMongoDB) GetAllFarms(skip int, perPage int, filters FarmFilters) ([]farm.FarmDto, error) {
	client, err := database.InitConnection()
	if err != nil {
		shared.LoggerError("Error initializing database connection", err)
		return []farm.FarmDto{}, err
	}
	shared.LoggerInfo("Database connection established")

	filter := bson.M{}

	if filters.Name != "" {
		safeName := regexp.QuoteMeta(filters.Name)
		filter["name"] = bson.M{"$regex": "^" + safeName, "$options": "i"}
	}
	if filters.Unit != "" {
		filter["unit"] = filters.Unit
	}
	if filters.CropType != "" {
		filter["crops.crop_type"] = filters.CropType
	}

	if filters.LandAreaLimit != "" && filters.LandAreaInit == "" {
		landAreaLimitFloat, err := strconv.ParseFloat(filters.LandAreaLimit, 64)
		if err != nil {
			return []farm.FarmDto{}, err
		}
		filter["land_area"] = bson.M{"$lt": landAreaLimitFloat}
	}

	if filters.LandAreaInit != "" && filters.LandAreaLimit == "" {
		landAreaInitFloat, err := strconv.ParseFloat(filters.LandAreaInit, 64)
		if err != nil {
			return []farm.FarmDto{}, err
		}
		filter["land_area"] = bson.M{"$gt": landAreaInitFloat}
	}

	if filters.LandAreaLimit != "" && filters.LandAreaInit != "" {
		landAreaLimitFloat, err := strconv.ParseFloat(filters.LandAreaLimit, 64)
		if err != nil {
			return []farm.FarmDto{}, err
		}
		landAreaInitFloat, err := strconv.ParseFloat(filters.LandAreaInit, 64)
		if err != nil {
			return []farm.FarmDto{}, err
		}
		filter["land_area"] = bson.M{"$lt": landAreaLimitFloat, "$gt": landAreaInitFloat}
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(perPage))
	findOptions.SetSort(bson.M{"name": 1})

	cursor, err := client.Collection(shared.FARM_SCHEMA).Find(context.Background(), filter, findOptions)
	if err != nil {
		shared.LoggerError("Error executing MongoDB query", err)
		return []farm.FarmDto{}, err
	}
	defer cursor.Close(context.Background())

	var farmData []farm.FarmDto
	if err := cursor.All(context.Background(), &farmData); err != nil {
		shared.LoggerError("Error retrieving farm data from cursor", err)
		return []farm.FarmDto{}, err
	}

	shared.LoggerInfo("Farm data successfully retrieved", zap.Int("totalFarms", len(farmData)))

	return farmData, nil
}

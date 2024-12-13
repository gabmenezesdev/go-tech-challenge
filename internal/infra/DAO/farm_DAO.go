package dao

import (
	"context"

	"github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
	"github.com/gabmenezesdev/go-tech-challenge/internal/infra/database"
	shared "github.com/gabmenezesdev/go-tech-challenge/internal/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Using the CQRS pattern, we create a separate query as it is more optimized for finding data.
// Since the query doesn't manipulate a domain (i.e., it doesn't receive, process, or modify entity), a DAO is more appropriate.

type FarmDaoMongoDB struct{}

func (f FarmDaoMongoDB) GetAllFarms(skip int, perPage int, filters FarmFilters) ([]farm.FarmDto, error) {
	client, err := database.InitConnection()
	if err != nil {
		return []farm.FarmDto{}, err
	}

	// if filters.Name != "" {
	// 	query += " AND name = ?"
	// }
	// if filters.LandArea != "" {
	// 	query += " AND location = ?"
	// }
	// if filters.Unit != "" {
	// 	query += " AND size >= ?"
	// }
	// if filters.Address != "" {
	// 	query += " AND size <= ?"
	// }
	// if filters.CropType != "" {
	// 	query += " AND is_organic = true"
	// }
	// if filters.IsIrrigated == true {
	// 	query += " AND is_organic = true"
	// }
	// if filters.IsInsured == true {
	// 	query += " AND is_organic = true"
	// }

	filter := bson.M{}

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(perPage))

	cursor, err := client.Collection(shared.FARM_SCHEMA).Find(context.Background(), filter, findOptions)
	if err != nil {
		return []farm.FarmDto{}, err
	}
	defer cursor.Close(context.Background())

	var farmData []farm.FarmDto
	if err := cursor.All(context.Background(), &farmData); err != nil {
		return []farm.FarmDto{}, err
	}

	return farmData, nil
}

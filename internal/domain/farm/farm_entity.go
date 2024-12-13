package farm

import (
	"errors"
	"time"

	crop "github.com/gabmenezesdev/go-tech-challenge/internal/domain/crop"
)

type Farm struct {
	id       int64
	name     FarmName
	landArea LandArea
	unit     UnitOfMeasure
	address  string
	crops    []string
}

type FarmDto struct {
	ID       string         `json:"_id" bson:"_id"`
	Name     string         `json:"name" bson:"name"`
	LandArea float64        `json:"land_area" bson:"land_area"`
	Unit     string         `json:"unit" bson:"unit"`
	Address  string         `json:"address" bson:"address"`
	Crops    []crop.CropDto `json:"crops" bson:"crops"`
}

func createFarm(name string, landArea float64, unit, address string, crops []string) (*Farm, error) {
	farmName, err := NewFarmName(name)
	if err != nil {
		return nil, err
	}

	unitOfMeasure, err := newUnitOfMeasure(unit)
	if err != nil {
		return nil, err
	}

	newLandArea, err := NewLandArea(landArea)
	if err != nil {
		return nil, err
	}

	if len(address) == 0 {
		return nil, errors.New("INVALID_ADDRESS")
	}

	return &Farm{
		name:     farmName,
		landArea: newLandArea,
		unit:     unitOfMeasure,
		address:  address,
		crops:    crops,
	}, nil
}

func NewFarm(name string, landArea float64, unit, address string) (*Farm, error) {
	return createFarm(name, landArea, unit, address, []string{})
}

func NewFarmWithID(farmId int64, name string, landArea float64, unit, address string, crops []string, createdAt time.Time) (*Farm, error) {
	farm, err := createFarm(name, landArea, unit, address, crops)
	if err != nil {
		return nil, err
	}
	farm.id = farmId
	return farm, nil
}

func (f *Farm) GetID() int64 {
	return f.id
}

func (f *Farm) GetName() string {
	return f.name.Value()
}

func (f *Farm) GetLandArea() float64 {
	return f.landArea.Value()
}

func (f *Farm) GetUnit() string {
	return f.unit.Value()
}

func (f *Farm) GetAddress() string {
	return f.address
}

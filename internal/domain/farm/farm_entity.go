package farm

import (
	"errors"
	"time"

	crop "github.com/gabmenezesdev/go-tech-challenge/internal/domain/crop"
	"github.com/gabmenezesdev/go-tech-challenge/internal/shared"
)

type Farm struct {
	id       string
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
	shared.LoggerInfo("Init create farm instance")
	farmName, err := NewFarmName(name)
	if err != nil {
		shared.LoggerError("Error while instantiating farm name", err)
		return nil, err
	}

	unitOfMeasure, err := newUnitOfMeasure(unit)
	if err != nil {
		shared.LoggerError("Error while instantiating unit of measure", err)
		return nil, err
	}

	newLandArea, err := NewLandArea(landArea)
	if err != nil {
		shared.LoggerError("Error while instantiating land area", err)
		return nil, err
	}

	if len(address) == 0 {
		shared.LoggerError("Invalid address provided", errors.New("INVALID_ADDRESS"))
		return nil, errors.New("INVALID_ADDRESS")
	}

	shared.LoggerInfo("Farm instance created successfully")
	return &Farm{
		name:     farmName,
		landArea: newLandArea,
		unit:     unitOfMeasure,
		address:  address,
		crops:    crops,
	}, nil
}

func NewFarm(name string, landArea float64, unit string, address string) (*Farm, error) {
	shared.LoggerInfo("Creating new farm without crops")
	return createFarm(name, landArea, unit, address, []string{})
}

func NewFarmWithID(farmId string, name string, landArea float64, unit, address string, crops []string, createdAt time.Time) (*Farm, error) {
	farm, err := createFarm(name, landArea, unit, address, crops)
	if err != nil {
		shared.LoggerError("Error while creating farm with ID", err)
		return nil, err
	}
	farm.SetId(farmId)
	shared.LoggerInfo("Farm with ID created successfully")
	return farm, nil
}

func (f *Farm) GetID() string {
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

func (f *Farm) SetId(id string) {
	f.id = id
}

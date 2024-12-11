package farm

import (
	"errors"
	"time"
)

type Farm struct {
	id        int64
	name      FarmName
	landArea  LandArea
	unit      UnitOfMeasure
	address   string
	crops     []string
	createdAt time.Time
}

func createFarm(name string, landArea float64, unit, address string, crops []string, createdAt time.Time) (*Farm, error) {
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
		return nil, errors.New("Invalid Address")
	}

	return &Farm{
		name:      farmName,
		landArea:  newLandArea,
		unit:      unitOfMeasure,
		address:   address,
		crops:     crops,
		createdAt: createdAt,
	}, nil
}

func NewFarm(name string, landArea float64, unit, address string, crops []string) (*Farm, error) {
	return createFarm(name, landArea, unit, address, crops, time.Now())
}

func NewFarmWithID(farmId int64, name string, landArea float64, unit, address string, crops []string, createdAt time.Time) (*Farm, error) {
	farm, err := createFarm(name, landArea, unit, address, crops, createdAt)
	if err != nil {
		return nil, err
	}
	farm.id = farmId
	return farm, nil
}

func (f *Farm) GetID() int64 {
	return f.id
}

func (f *Farm) GetCreatedAt() time.Time {
	return f.createdAt
}

func (f *Farm) Name() string {
	return f.name.Value()
}

package farm

import (
	"errors"
	"time"
)

type Farm struct {
	id       int64
	name     FarmName
	landArea LandArea
	unit     UnitOfMeasure
	address  string
	crops    []string
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
		return nil, errors.New("Invalid Address")
	}

	return &Farm{
		name:     farmName,
		landArea: newLandArea,
		unit:     unitOfMeasure,
		address:  address,
		crops:    crops,
	}, nil
}

func NewFarm(name string, landArea float64, unit, address string, crops []string) (*Farm, error) {
	return createFarm(name, landArea, unit, address, crops)
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

package dao

import (
	"github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
)

type FarmFilters struct {
	Name        string
	Unit        string
	CropType    string
	IsIrrigated string
	IsInsured   string
}

type FarmDao interface {
	GetAllFarms(skip int, perPage int, filters FarmFilters) ([]farm.FarmDto, error)
}

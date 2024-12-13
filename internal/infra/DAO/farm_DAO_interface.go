package dao

import (
	"github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"
)

type FarmFilters struct {
	Name        string
	LandArea    string
	Unit        string
	Address     string
	CropType    string
	IsIrrigated bool
	IsInsured   bool
}

type FarmDao interface {
	GetAllFarms(skip int, perPage int, filters FarmFilters) ([]farm.FarmDto, error)
}
package dao

import "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"

type FarmDao interface {
	GetAllFarms(skip int, perPage int) ([]farm.FarmDto, error)
}

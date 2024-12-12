package repository

import farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"

type FarmRepository interface {
	CreateFarm(u *farm.Farm) (string, error)
	DeleteFarmById(farmId string) error
	GetFarmById(farmId string) (farm.Farm, error)
}

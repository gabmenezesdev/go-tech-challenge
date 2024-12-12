package repository

import farm "github.com/gabmenezesdev/go-tech-challenge/internal/domain/farm"

type FarmRepository interface {
	CreateFarm(u *farm.Farm) (string, error)
	DeleteFarmById(id int64) error
	GetFarmById(id int64) error
}

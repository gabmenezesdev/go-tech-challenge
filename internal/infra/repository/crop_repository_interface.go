package repository

import crop "github.com/gabmenezesdev/go-tech-challenge/internal/domain/crop"

type CropRepository interface {
	CreateCrop(crop *crop.Crop, farmId string) error
	DeleteCropById(id int64) error
}

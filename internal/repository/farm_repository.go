package repository

import (
	"gorm.io/gorm"
)

type FarmRepository interface {
}

type farmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

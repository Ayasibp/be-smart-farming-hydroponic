package repository

import (
	"gorm.io/gorm"
)

type GrowthHistRepository interface {
	
}

type growthHistRepository struct {
	db *gorm.DB
}

func NewGrowthHistRepository(db *gorm.DB) GrowthHistRepository {
	return &growthHistRepository{
		db: db,
	}
}

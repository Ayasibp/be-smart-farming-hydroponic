package repository

import (
	"gorm.io/gorm"
)

type AggregationRepository interface {
}

type aggregationRepository struct {
	db *gorm.DB
}

func NewAggregationRepository(db *gorm.DB) AggregationRepository {
	return &aggregationRepository{
		db: db,
	}
}

package repository

import (
	"gorm.io/gorm"
)

type SystemUnitRepository interface {
	
}

type systemUnitRepository struct {
	db *gorm.DB
}

func NewSystemUnitRepository(db *gorm.DB) SystemUnitRepository {
	return &systemUnitRepository{
		db: db,
	}
}
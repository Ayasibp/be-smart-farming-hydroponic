package repository

import (
	"gorm.io/gorm"
)

type TankTransRepository interface {
	
}

type tankTransRepository struct {
	db *gorm.DB
}

func NewTankTransRepository(db *gorm.DB) TankTransRepository {
	return &tankTransRepository{
		db: db,
	}
}


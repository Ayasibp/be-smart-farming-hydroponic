package repository

import (
	"gorm.io/gorm"
)

type UnitIdRepository interface {
}
type unitIdRepository struct {
	db *gorm.DB
}

func NewUnitIdRepository(db *gorm.DB) UnitIdRepository {
	return &unitIdRepository{
		db: db,
	}
}

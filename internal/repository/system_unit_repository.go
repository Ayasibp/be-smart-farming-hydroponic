package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type SystemUnitRepository interface {
	CreateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error)
}

type systemUnitRepository struct {
	db *gorm.DB
}

func NewSystemUnitRepository(db *gorm.DB) SystemUnitRepository {
	return &systemUnitRepository{
		db: db,
	}
}

func (r systemUnitRepository) CreateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error) {
	res := r.db.Raw("INSERT INTO hydroponic_system.system_units (farm_id ,unit_key, tank_volume , tank_a_volume , tank_b_volume , created_at) VALUES (?,?,?,?,?,?) RETURNING *;", inputModel.FarmId, inputModel.UnitKey, inputModel.TankVolume, inputModel.TankAVolume, inputModel.TankBVolume, time.Now()).Scan(inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}

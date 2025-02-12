package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type TankTransRepository interface {
	CreateTankTransaction(inputModel *model.TankTran) (*model.TankTran, error)
}

type tankTransRepository struct {
	db *gorm.DB
}

func NewTankTransRepository(db *gorm.DB) TankTransRepository {
	return &tankTransRepository{
		db: db,
	}
}

func (r *tankTransRepository) CreateTankTransaction(inputModel *model.TankTran) (*model.TankTran, error) {

	sqlScript := `INSERT INTO hydroponic_system.tank_trans(farm_id, system_id, water_volume, a_volume,b_volume, created_at) 
				VALUES (?,?,?,?,?,?) 
				RETURNING id, farm_id, system_id, water_volume, a_volume, b_volume;`

	res := r.db.Raw(sqlScript, inputModel.FarmId, inputModel.SystemId, inputModel.WaterVolume, inputModel.AVolume, inputModel.BVolume, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}

package repository

import (
	"time"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type SystemUnitRepository interface {
	CreateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error)
	UpdateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error)
	GetSystemUnits(farmsId *string) ([]*model.SystemUnitJoined, error)
	GetSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error)
	DeleteSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error)
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

	sqlScript := `INSERT INTO hydroponic_system.system_units (farm_id ,unit_key, tank_volume , tank_a_volume , tank_b_volume , created_at) 
				VALUES (?,?,?,?,?,?) 
				RETURNING farm_id, unit_key, tank_volume, tank_a_volume, tank_b_volume;`

	res := r.db.Raw(sqlScript,
		inputModel.FarmId,
		inputModel.UnitKey,
		inputModel.TankVolume,
		inputModel.TankAVolume,
		inputModel.TankBVolume,
		time.Now()).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}

func (r systemUnitRepository) GetSystemUnits(farmsId *string) ([]*model.SystemUnitJoined, error) {
	var inputModel []*model.SystemUnitJoined
	sqlScript := `SELECT su.id,su.unit_key, su.farm_id, f."name" as farm_name, su.tank_volume, su.tank_a_volume , su.tank_b_volume
				FROM hydroponic_system.system_units su 
				LEFT JOIN hydroponic_system.farms f on f.id = su.farm_id 
				WHERE su.deleted_at is NULL 
				` + *farmsId

	res := r.db.Raw(sqlScript).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}

func (r systemUnitRepository) GetSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error) {

	sqlScript := `SELECT id, farm_id, unit_key, tank_volume, tank_a_volume, tank_b_volume 
				FROM hydroponic_system.system_units 
				WHERE id = ?`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidSystemUnitID
	}
	return inputModel, nil

}

func (r systemUnitRepository) UpdateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error) {

	sqlScript := `UPDATE hydroponic_system.system_units 
				SET updated_at = ?, unit_key = ?, farm_id = ?,tank_volume = ?, tank_a_volume = ?, tank_b_volume = ? 
				WHERE id = ? 
				RETURNING id,farm_id, unit_key, tank_volume, tank_a_volume, tank_b_volume `

	res := r.db.Raw(sqlScript,
		time.Now(),
		inputModel.UnitKey,
		inputModel.FarmId,
		inputModel.TankVolume,
		inputModel.TankAVolume,
		inputModel.TankBVolume,
		inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidProfileID
	}
	return inputModel, nil
}

func (r systemUnitRepository) DeleteSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error) {

	sqlScript := `UPDATE hydroponic_system.system_units 
				SET deleted_at = ? 
				WHERE id = ? 
				RETURNING id`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {

		return nil, errs.InvalidSystemUnitID
	}
	return inputModel, nil
}

package repository

import (
	"time"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"gorm.io/gorm"
)

type SystemUnitRepository interface {
	CreateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error)
	UpdateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error)
	GetSystemUnits(farmId *string) ([]*model.SystemUnitJoined, error)
	GetSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error)
	DeleteSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error)
}

type systemUnitRepository struct {
	db *gorm.DB
}

func NewSystemUnitRepository(db *gorm.DB) SystemUnitRepository {
	return &systemUnitRepository{db: db}
}

func (r *systemUnitRepository) CreateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error) {
	logger.Info("systemUnitRepository", "Creating a new system unit", "farmId", inputModel.FarmId, "unitKey", inputModel.UnitKey)

	sqlScript := `INSERT INTO hydroponic_system.system_units (farm_id, unit_key, tank_volume, tank_a_volume, tank_b_volume, created_at) 
				  VALUES (?, ?, ?, ?, ?, ?) 
				  RETURNING id, farm_id, unit_key, tank_volume, tank_a_volume, tank_b_volume;`

	res := r.db.Raw(sqlScript,
		inputModel.FarmId,
		inputModel.UnitKey,
		inputModel.TankVolume,
		inputModel.TankAVolume,
		inputModel.TankBVolume,
		time.Now()).Scan(inputModel)

	if res.Error != nil {
		logger.Error("systemUnitRepository", "Failed to create system unit", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("systemUnitRepository", "System unit created successfully", "id", inputModel.ID, "unitKey", inputModel.UnitKey)
	return inputModel, nil
}

func (r *systemUnitRepository) GetSystemUnits(farmId *string) ([]*model.SystemUnitJoined, error) {
	logger.Info("systemUnitRepository", "Fetching system units", "farmId", *farmId)

	var units []*model.SystemUnitJoined
	sqlScript := `SELECT su.id, su.unit_key, su.farm_id, f.name as farm_name, su.tank_volume, su.tank_a_volume, su.tank_b_volume
				  FROM hydroponic_system.system_units su
				  LEFT JOIN hydroponic_system.farms f ON f.id = su.farm_id
				  WHERE su.deleted_at IS NULL AND su.farm_id = ?`

	res := r.db.Raw(sqlScript, *farmId).Scan(&units)

	if res.Error != nil {
		logger.Error("systemUnitRepository", "Failed to fetch system units", "farmId", *farmId, "error", res.Error)
		return nil, res.Error
	}

	logger.Info("systemUnitRepository", "Successfully fetched system units", "count", len(units))
	return units, nil
}

func (r *systemUnitRepository) GetSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error) {
	logger.Info("systemUnitRepository", "Fetching system unit by ID", "id", inputModel.ID)

	sqlScript := `SELECT id, farm_id, unit_key, tank_volume, tank_a_volume, tank_b_volume
				  FROM hydroponic_system.system_units
				  WHERE id = ?`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(inputModel)

	if res.Error != nil {
		logger.Error("systemUnitRepository", "Failed to fetch system unit", "id", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("systemUnitRepository", "No system unit found", "id", inputModel.ID)
		return nil, errs.InvalidSystemUnitID
	}

	logger.Info("systemUnitRepository", "Successfully fetched system unit", "id", inputModel.ID)
	return inputModel, nil
}

func (r *systemUnitRepository) UpdateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error) {
	logger.Info("systemUnitRepository", "Updating system unit", "id", inputModel.ID)

	sqlScript := `UPDATE hydroponic_system.system_units 
				  SET updated_at = ?, unit_key = ?, farm_id = ?, tank_volume = ?, tank_a_volume = ?, tank_b_volume = ? 
				  WHERE id = ? 
				  RETURNING id, farm_id, unit_key, tank_volume, tank_a_volume, tank_b_volume`

	res := r.db.Raw(sqlScript,
		time.Now(),
		inputModel.UnitKey,
		inputModel.FarmId,
		inputModel.TankVolume,
		inputModel.TankAVolume,
		inputModel.TankBVolume,
		inputModel.ID).Scan(inputModel)

	if res.Error != nil {
		logger.Error("systemUnitRepository", "Failed to update system unit", "id", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("systemUnitRepository", "No system unit updated", "id", inputModel.ID)
		return nil, errs.InvalidSystemUnitID
	}

	logger.Info("systemUnitRepository", "Successfully updated system unit", "id", inputModel.ID)
	return inputModel, nil
}

func (r *systemUnitRepository) DeleteSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error) {
	logger.Info("systemUnitRepository", "Deleting system unit", "id", inputModel.ID)

	sqlScript := `UPDATE hydroponic_system.system_units 
				  SET deleted_at = ? 
				  WHERE id = ? 
				  RETURNING id`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(inputModel)

	if res.Error != nil {
		logger.Error("systemUnitRepository", "Failed to delete system unit", "id", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("systemUnitRepository", "No system unit found to delete", "id", inputModel.ID)
		return nil, errs.InvalidSystemUnitID
	}

	logger.Info("systemUnitRepository", "Successfully deleted system unit", "id", inputModel.ID)
	return inputModel, nil
}

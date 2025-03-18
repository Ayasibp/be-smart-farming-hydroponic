package repository

import (
	"strconv"
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
	logger.Info("systemUnitRepository", "Creating a new system unit", map[string]string{
		"farmId":  inputModel.FarmId.String(),
		"unitKey": inputModel.UnitKey.String(),
	})

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
		logger.Error("systemUnitRepository", "Failed to create system unit", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("systemUnitRepository", "System unit created successfully", map[string]string{
		"id":      inputModel.ID.String(),
		"unitKey": inputModel.UnitKey.String(),
	})
	return inputModel, nil
}

func (r *systemUnitRepository) GetSystemUnits(farmId *string) ([]*model.SystemUnitJoined, error) {
	logger.Info("systemUnitRepository", "Fetching system units", map[string]string{
		"farmId": *farmId,
	})

	var units []*model.SystemUnitJoined
	sqlScript := `SELECT su.id, su.unit_key, su.farm_id, f.name as farm_name, su.tank_volume, su.tank_a_volume, su.tank_b_volume
				  FROM hydroponic_system.system_units su
				  LEFT JOIN hydroponic_system.farms f ON f.id = su.farm_id
				  WHERE su.deleted_at IS NULL AND su.farm_id = ?`

	res := r.db.Raw(sqlScript, *farmId).Scan(&units)

	if res.Error != nil {
		logger.Error("systemUnitRepository", "Failed to fetch system units", map[string]string{
			"farmId": *farmId,
			"error":  res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("systemUnitRepository", "Successfully fetched system units", map[string]string{
		"count": strconv.Itoa(len(units)),
	})
	return units, nil
}

func (r *systemUnitRepository) GetSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error) {
	logger.Info("systemUnitRepository", "Fetching system unit by ID", map[string]string{
		"id": inputModel.ID.String(),
	})

	sqlScript := `SELECT id, farm_id, unit_key, tank_volume, tank_a_volume, tank_b_volume
				  FROM hydroponic_system.system_units
				  WHERE id = ?`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(inputModel)

	if res.Error != nil {
		logger.Error("systemUnitRepository", "Failed to fetch system unit", map[string]string{
			"id":    inputModel.ID.String(),
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("systemUnitRepository", "No system unit found", map[string]string{
			"id": inputModel.ID.String(),
		})
		return nil, errs.InvalidSystemUnitID
	}

	logger.Info("systemUnitRepository", "Successfully fetched system unit", map[string]string{
		"id": inputModel.ID.String(),
	})
	return inputModel, nil
}

func (r *systemUnitRepository) UpdateSystemUnit(inputModel *model.SystemUnit) (*model.SystemUnit, error) {
	logger.Info("systemUnitRepository", "Updating system unit", map[string]string{
		"id": inputModel.ID.String(),
	})

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
		logger.Error("systemUnitRepository", "Failed to update system unit", map[string]string{
			"id":    inputModel.ID.String(),
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("systemUnitRepository", "No system unit updated", map[string]string{
			"id": inputModel.ID.String(),
		})
		return nil, errs.InvalidSystemUnitID
	}

	logger.Info("systemUnitRepository", "Successfully updated system unit", map[string]string{
		"id": inputModel.ID.String(),
	})
	return inputModel, nil
}

func (r *systemUnitRepository) DeleteSystemUnitById(inputModel *model.SystemUnit) (*model.SystemUnit, error) {
	logger.Info("systemUnitRepository", "Deleting system unit", map[string]string{
		"id": inputModel.ID.String(),
	})

	sqlScript := `UPDATE hydroponic_system.system_units 
				  SET deleted_at = ? 
				  WHERE id = ? 
				  RETURNING id`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(inputModel)

	if res.Error != nil {
		logger.Error("systemUnitRepository", "Failed to delete system unit", map[string]string{
			"id":    inputModel.ID.String(),
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("systemUnitRepository", "No system unit found to delete", map[string]string{
			"id": inputModel.ID.String(),
		})
		return nil, errs.InvalidSystemUnitID
	}

	logger.Info("systemUnitRepository", "Successfully deleted system unit", map[string]string{
		"id": inputModel.ID.String(),
	})
	return inputModel, nil
}

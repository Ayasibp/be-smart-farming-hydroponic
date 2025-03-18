package repository

import (
	"strconv"
	"time"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"gorm.io/gorm"
)

type FarmRepository interface {
	CreateFarm(inputModel *model.Farm) (*model.Farm, error)
	GetFarms() ([]*model.Farm, error)
	GetFarmById(inputModel *model.Farm) (*model.Farm, error)
	UpdateFarm(inputModel *model.Farm) (*model.Farm, error)
	DeleteFarm(inputModel *model.Farm) (*model.Farm, error)
}

type farmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) FarmRepository {
	return &farmRepository{
		db: db,
	}
}

func (r *farmRepository) CreateFarm(inputModel *model.Farm) (*model.Farm, error) {
	logger.Info("farmRepository", "Creating farm", map[string]string{
		"profileID": inputModel.ProfileId.String(),
		"name":      inputModel.Name,
	})

	sqlScript := `INSERT INTO hydroponic_system.farms (profile_id , name , address, created_at) 
				VALUES (?,?,?,?) 
				RETURNING profile_id, name, address;`

	res := r.db.Raw(sqlScript, inputModel.ProfileId, inputModel.Name, inputModel.Address, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to create farm", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("farmRepository", "Farm created successfully", map[string]string{
		"profileID": inputModel.ProfileId.String(),
		"name":      inputModel.Name,
	})
	return inputModel, nil
}

func (r *farmRepository) GetFarms() ([]*model.Farm, error) {
	logger.Info("farmRepository", "Fetching all farms", nil)

	var farms []*model.Farm

	sqlScript := `SELECT * FROM hydroponic_system.farms 
				WHERE deleted_at IS NULL`

	res := r.db.Raw(sqlScript).Scan(&farms)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to fetch farms", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("farmRepository", "Fetched farms successfully", map[string]string{
		"count": strconv.Itoa(len(farms)),
	})
	return farms, nil
}

func (r *farmRepository) GetFarmById(inputModel *model.Farm) (*model.Farm, error) {
	logger.Info("farmRepository", "Fetching farm by ID", map[string]string{
		"farmID": inputModel.ID.String(),
	})

	sqlScript := `SELECT * FROM hydroponic_system.farms 
				WHERE id = ?`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to fetch farm", map[string]string{
			"farmID": inputModel.ID.String(),
			"error":  res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("farmRepository", "Farm ID not found", map[string]string{
			"farmID": inputModel.ID.String(),
		})
		return nil, errs.InvalidFarmID
	}

	logger.Info("farmRepository", "Farm fetched successfully", map[string]string{
		"farmID": inputModel.ID.String(),
	})
	return inputModel, nil
}

func (r *farmRepository) UpdateFarm(inputModel *model.Farm) (*model.Farm, error) {
	logger.Info("farmRepository", "Updating farm", map[string]string{
		"farmID": inputModel.ID.String(),
		"name":   inputModel.Name,
	})

	sqlScript := `UPDATE hydroponic_system.farms 
				SET updated_at = ?, name = ?, address = ?  
				WHERE id = ? 
				RETURNING *`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.Name, inputModel.Address, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to update farm", map[string]string{
			"farmID": inputModel.ID.String(),
			"error":  res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("farmRepository", "Farm ID not found for update", map[string]string{
			"farmID": inputModel.ID.String(),
		})
		return nil, errs.InvalidFarmID
	}

	logger.Info("farmRepository", "Farm updated successfully", map[string]string{
		"farmID": inputModel.ID.String(),
		"name":   inputModel.Name,
	})
	return inputModel, nil
}

func (r *farmRepository) DeleteFarm(inputModel *model.Farm) (*model.Farm, error) {
	logger.Info("farmRepository", "Deleting farm", map[string]string{
		"farmID": inputModel.ID.String(),
	})

	sqlScript := `UPDATE hydroponic_system.farms 
				SET deleted_at = ? 
				WHERE id = ? 
				RETURNING *`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to delete farm", map[string]string{
			"farmID": inputModel.ID.String(),
			"error":  res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("farmRepository", "Farm ID not found for deletion", map[string]string{
			"farmID": inputModel.ID.String(),
		})
		return nil, errs.InvalidFarmID
	}

	logger.Info("farmRepository", "Farm deleted successfully", map[string]string{
		"farmID": inputModel.ID.String(),
	})
	return inputModel, nil
}

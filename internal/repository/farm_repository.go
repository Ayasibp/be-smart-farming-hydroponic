package repository

import (
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
	logger.Info("farmRepository", "Creating farm", "profileID", inputModel.ProfileId, "name", inputModel.Name)

	sqlScript := `INSERT INTO hydroponic_system.farms (profile_id , name , address, created_at) 
				VALUES (?,?,?,?) 
				RETURNING profile_id, name, address;`

	res := r.db.Raw(sqlScript, inputModel.ProfileId, inputModel.Name, inputModel.Address, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to create farm", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("farmRepository", "Farm created successfully", "profileID", inputModel.ProfileId, "name", inputModel.Name)
	return inputModel, nil
}

func (r *farmRepository) GetFarms() ([]*model.Farm, error) {
	logger.Info("farmRepository", "Fetching all farms")

	var farms []*model.Farm

	sqlScript := `SELECT * FROM hydroponic_system.farms 
				WHERE deleted_at IS NULL`

	res := r.db.Raw(sqlScript).Scan(&farms)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to fetch farms", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("farmRepository", "Fetched farms successfully", "count", len(farms))
	return farms, nil
}

func (r *farmRepository) GetFarmById(inputModel *model.Farm) (*model.Farm, error) {
	logger.Info("farmRepository", "Fetching farm by ID", "farmID", inputModel.ID)

	sqlScript := `SELECT * FROM hydroponic_system.farms 
				WHERE id = ?`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to fetch farm", "farmID", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("farmRepository", "Farm ID not found", "farmID", inputModel.ID)
		return nil, errs.InvalidFarmID
	}

	logger.Info("farmRepository", "Farm fetched successfully", "farmID", inputModel.ID)
	return inputModel, nil
}

func (r *farmRepository) UpdateFarm(inputModel *model.Farm) (*model.Farm, error) {
	logger.Info("farmRepository", "Updating farm", "farmID", inputModel.ID, "name", inputModel.Name)

	sqlScript := `UPDATE hydroponic_system.farms 
				SET updated_at = ?, name = ?, address = ?  
				WHERE id = ? 
				RETURNING *`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.Name, inputModel.Address, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to update farm", "farmID", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("farmRepository", "Farm ID not found for update", "farmID", inputModel.ID)
		return nil, errs.InvalidFarmID
	}

	logger.Info("farmRepository", "Farm updated successfully", "farmID", inputModel.ID, "name", inputModel.Name)
	return inputModel, nil
}

func (r *farmRepository) DeleteFarm(inputModel *model.Farm) (*model.Farm, error) {
	logger.Info("farmRepository", "Deleting farm", "farmID", inputModel.ID)

	sqlScript := `UPDATE hydroponic_system.farms 
				SET deleted_at = ? 
				WHERE id = ? 
				RETURNING *`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("farmRepository", "Failed to delete farm", "farmID", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("farmRepository", "Farm ID not found for deletion", "farmID", inputModel.ID)
		return nil, errs.InvalidFarmID
	}

	logger.Info("farmRepository", "Farm deleted successfully", "farmID", inputModel.ID)
	return inputModel, nil
}

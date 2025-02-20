package repository

import (
	"time"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"gorm.io/gorm"
)

type UnitIdRepository interface {
	CreateUnitId() (*model.UnitId, error)
	GetUnitIds() ([]*model.UnitId, error)
	GetUnitIdById(inputModel *model.UnitId) (*model.UnitId, error)
	DeleteUnitIdById(inputModel *model.UnitId) (*model.UnitId, error)
}

type unitIdRepository struct {
	db *gorm.DB
}

func NewUnitIdRepository(db *gorm.DB) UnitIdRepository {
	return &unitIdRepository{db: db}
}

func (r *unitIdRepository) CreateUnitId() (*model.UnitId, error) {
	logger.Info("unitIdRepository", "Creating a new unit ID")

	inputModel := &model.UnitId{}

	sqlScript := `INSERT INTO super_admin.unit_ids (created_at) 
				  VALUES (?) 
				  RETURNING id, created_at;`

	res := r.db.Raw(sqlScript, time.Now()).Scan(inputModel)

	if res.Error != nil {
		logger.Error("unitIdRepository", "Failed to create unit ID", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("unitIdRepository", "Successfully created unit ID", "id", inputModel.ID)
	return inputModel, nil
}

func (r *unitIdRepository) GetUnitIds() ([]*model.UnitId, error) {
	logger.Info("unitIdRepository", "Fetching all unit IDs")

	var unitIds []*model.UnitId

	sqlScript := `SELECT id FROM super_admin.unit_ids 
				  WHERE deleted_at IS NULL;`

	res := r.db.Raw(sqlScript).Scan(&unitIds)

	if res.Error != nil {
		logger.Error("unitIdRepository", "Failed to fetch unit IDs", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("unitIdRepository", "Successfully fetched unit IDs", "count", len(unitIds))
	return unitIds, nil
}

func (r *unitIdRepository) GetUnitIdById(inputModel *model.UnitId) (*model.UnitId, error) {
	logger.Info("unitIdRepository", "Fetching unit ID", "id", inputModel.ID)

	sqlScript := `SELECT id FROM super_admin.unit_ids 
				  WHERE id = ? AND deleted_at IS NULL;`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(inputModel)

	if res.Error != nil {
		logger.Error("unitIdRepository", "Failed to fetch unit ID", "id", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("unitIdRepository", "Unit ID not found", "id", inputModel.ID)
		return nil, errs.InvalidUnitKey
	}

	logger.Info("unitIdRepository", "Successfully fetched unit ID", "id", inputModel.ID)
	return inputModel, nil
}

func (r *unitIdRepository) DeleteUnitIdById(inputModel *model.UnitId) (*model.UnitId, error) {
	logger.Info("unitIdRepository", "Deleting unit ID", "id", inputModel.ID)

	sqlScript := `UPDATE super_admin.unit_ids 
				  SET deleted_at = ? 
				  WHERE id = ? 
				  RETURNING id;`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(inputModel)

	if res.Error != nil {
		logger.Error("unitIdRepository", "Failed to delete unit ID", "id", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("unitIdRepository", "Unit ID not found for deletion", "id", inputModel.ID)
		return nil, errs.InvalidUnitKey
	}

	logger.Info("unitIdRepository", "Successfully deleted unit ID", "id", inputModel.ID)
	return inputModel, nil
}

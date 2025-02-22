package repository

import (
	"strconv"
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
	logger.Info("unitIdRepository", "Creating a new unit ID", nil)

	inputModel := &model.UnitId{}

	sqlScript := `INSERT INTO super_admin.unit_ids (created_at) 
				  VALUES (?) 
				  RETURNING id, created_at;`

	res := r.db.Raw(sqlScript, time.Now()).Scan(inputModel)

	if res.Error != nil {
		logger.Error("unitIdRepository", "Failed to create unit ID", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("unitIdRepository", "Successfully created unit ID", map[string]string{
		"id": inputModel.ID.String(),
	})
	return inputModel, nil
}

func (r *unitIdRepository) GetUnitIds() ([]*model.UnitId, error) {
	logger.Info("unitIdRepository", "Fetching all unit IDs", nil)

	var unitIds []*model.UnitId

	sqlScript := `SELECT id FROM super_admin.unit_ids 
				  WHERE deleted_at IS NULL;`

	res := r.db.Raw(sqlScript).Scan(&unitIds)

	if res.Error != nil {
		logger.Error("unitIdRepository", "Failed to fetch unit IDs", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("unitIdRepository", "Successfully fetched unit IDs", map[string]string{
		"count": strconv.Itoa(len(unitIds)),
	})
	return unitIds, nil
}

func (r *unitIdRepository) GetUnitIdById(inputModel *model.UnitId) (*model.UnitId, error) {
	logger.Info("unitIdRepository", "Fetching unit ID", map[string]string{
		"id": inputModel.ID.String(),
	})

	sqlScript := `SELECT id FROM super_admin.unit_ids 
				  WHERE id = ? AND deleted_at IS NULL;`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(inputModel)

	if res.Error != nil {
		logger.Error("unitIdRepository", "Failed to fetch unit ID", map[string]string{
			"id":    inputModel.ID.String(),
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("unitIdRepository", "Unit ID not found", map[string]string{
			"id": inputModel.ID.String(),
		})
		return nil, errs.InvalidUnitKey
	}

	logger.Info("unitIdRepository", "Successfully fetched unit ID", map[string]string{
		"id": inputModel.ID.String(),
	})
	return inputModel, nil
}

func (r *unitIdRepository) DeleteUnitIdById(inputModel *model.UnitId) (*model.UnitId, error) {
	logger.Info("unitIdRepository", "Deleting unit ID", map[string]string{
		"id": inputModel.ID.String(),
	})

	sqlScript := `UPDATE super_admin.unit_ids 
				  SET deleted_at = ? 
				  WHERE id = ? 
				  RETURNING id;`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(inputModel)

	if res.Error != nil {
		logger.Error("unitIdRepository", "Failed to delete unit ID", map[string]string{
			"id":    inputModel.ID.String(),
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("unitIdRepository", "Unit ID not found for deletion", map[string]string{
			"id": inputModel.ID.String(),
		})
		return nil, errs.InvalidUnitKey
	}

	logger.Info("unitIdRepository", "Successfully deleted unit ID", map[string]string{
		"id": inputModel.ID.String(),
	})
	return inputModel, nil
}

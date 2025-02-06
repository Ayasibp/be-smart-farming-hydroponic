package repository

import (
	"time"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
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
	return &unitIdRepository{
		db: db,
	}
}

func (r unitIdRepository) CreateUnitId() (*model.UnitId, error) {
	var inputModel *model.UnitId

	sqlScript := `INSERT INTO super_admin.unit_ids (created_at) 
				VALUES (?) 
				RETURNING id;`

	res := r.db.Raw(sqlScript, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}
func (r unitIdRepository) GetUnitIds() ([]*model.UnitId, error) {
	var inputModel []*model.UnitId

	sqlScript := `SELECT id FROM super_admin.unit_ids 
				WHERE deleted_at IS NULL;`

	res := r.db.Raw(sqlScript).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}

func (r unitIdRepository) GetUnitIdById(inputModel *model.UnitId) (*model.UnitId, error) {

	sqlScript := `SELECT id FROM super_admin.unit_ids 
				WHERE id = ? AND deleted_at IS NULL;`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidUnitKey
	}
	return inputModel, nil

}

func (r unitIdRepository) DeleteUnitIdById(inputModel *model.UnitId) (*model.UnitId, error) {

	sqlScript := `UPDATE super_admin.unit_ids 
				SET deleted_at = ? 
				WHERE id = ? 
				RETURNING id`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {

		return nil, errs.InvalidUnitKey
	}
	return inputModel, nil
}

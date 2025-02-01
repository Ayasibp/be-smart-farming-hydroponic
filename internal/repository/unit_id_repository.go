package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type UnitIdRepository interface {
	CreateUnitId() (*model.UnitId, error)
	GetUnitIds() ([]*model.UnitId, error)
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

	res := r.db.Raw("INSERT INTO super_admin.unit_ids (created_at) VALUES (?) RETURNING *;", time.Now()).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}
func (r unitIdRepository) GetUnitIds() ([]*model.UnitId, error) {
	var inputModel []*model.UnitId

	res := r.db.Raw("SELECT * FROM super_admin.unit_ids;").Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}

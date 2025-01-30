package repository

import (
	"time"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type FarmRepository interface {
	CreateFarm(inputModel *model.Farm) (*model.Farm, error)
	GetFarms() ([]*model.Farm, error)
	GetFarmById(inputModel *model.Farm) (*model.Farm, error)
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

func (r farmRepository) CreateFarm(inputModel *model.Farm) (*model.Farm, error) {
	res := r.db.Raw("INSERT INTO farms (profile_id , name , address, created_at) VALUES (?,?,?,?) RETURNING *;", inputModel.ProfileId, inputModel.Name, inputModel.Address, time.Now()).Scan(inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}

func (r farmRepository) GetFarms() ([]*model.Farm, error) {

	var farms []*model.Farm

	res := r.db.Raw("SELECT * FROM farms WHERE deleted_at IS NULL").Scan(&farms)

	if res.Error != nil {
		return nil, res.Error
	}

	return farms, nil

}

func (r farmRepository) GetFarmById(inputModel *model.Farm) (*model.Farm, error) {

	res := r.db.Raw("SELECT * FROM farms WHERE id = ?", inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidFarmID
	}
	return inputModel, nil

}

func (r farmRepository) DeleteFarm(inputModel *model.Farm) (*model.Farm, error) {

	res := r.db.Raw("UPDATE farms SET deleted_at = ? WHERE id = ? RETURNING *", time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidFarmID
	}
	return inputModel, nil
}
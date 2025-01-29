package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type FarmRepository interface {
	CreateFarm(inputModel *model.Farm) (*model.Farm, error)
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

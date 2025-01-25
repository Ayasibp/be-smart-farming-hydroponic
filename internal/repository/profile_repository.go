package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	CreateProfile(input *dto.CreateProfile) (*model.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

func (r profileRepository) CreateProfile(input *dto.CreateProfile) (*model.Profile, error) {

	var inputModel = &model.Profile{
		ID:      input.AccountID,
		Name:    input.Name,
		Address: input.Address,
	}

	res := r.db.Raw("INSERT INTO profiles (account_id , name , address, created_at) VALUES (?,?,?,?) RETURNING *;", input.AccountID, input.Name, input.Address, time.Now()).Scan(inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}

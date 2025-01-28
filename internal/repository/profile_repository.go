package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	CreateProfile(input *dto.CreateProfile) (*model.Profile, error)
	DeleteProfile(id uuid.UUID) (*model.Profile, error)
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

func (r profileRepository) DeleteProfile(id uuid.UUID) (*model.Profile, error) {

	var inputModel *model.Profile

	res := r.db.Raw("Update profiles SET deleted_at = ? where ? RETURNING id", time.Now(),id).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}


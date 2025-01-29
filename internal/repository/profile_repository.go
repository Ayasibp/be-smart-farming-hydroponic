package repository

import (
	"time"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	CreateProfile(inputModel *model.Profile) (*model.Profile, error)
	GetProfileById(inputModel *model.Profile) (*model.Profile, error)
	DeleteProfile(inputModel *model.Profile) (*model.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

func (r profileRepository) CreateProfile(inputModel *model.Profile) (*model.Profile, error) {
	res := r.db.Raw("INSERT INTO profiles (account_id , name , address, created_at) VALUES (?,?,?,?) RETURNING *;", inputModel.AccountId, inputModel.Name, inputModel.Address, time.Now()).Scan(inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}

func (r profileRepository) GetProfileById(inputModel *model.Profile) (*model.Profile, error) {

	res := r.db.Raw("SELECT * FROM profiles WHERE id = ?", inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidProfileID
	}
	return inputModel, nil

}

func (r profileRepository) DeleteProfile(inputModel *model.Profile) (*model.Profile, error) {

	res := r.db.Raw("UPDATE profiles SET deleted_at = ? WHERE id = ? RETURNING *", time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidProfileID
	}
	return inputModel, nil

}

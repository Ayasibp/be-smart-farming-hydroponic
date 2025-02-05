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
	GetProfiles() ([]*model.Profile, error)
	UpdateProfile(inputModel *model.Profile) (*model.Profile, error)
	DeleteProfile(inputModel *model.Profile) (*model.Profile, error)
	CheckCreatedProfileByAccountId(inputModel *model.Profile) (*model.Profile, error)
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

	sqlScript := `INSERT INTO hydroponic_system.profiles (account_id , name , address, created_at) 
				VALUES (?,?,?,?) 
				RETURNING account_id, name, address;`

	res := r.db.Raw(sqlScript, inputModel.AccountId, inputModel.Name, inputModel.Address, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}

func (r profileRepository) GetProfiles() ([]*model.Profile, error) {

	var profiles []*model.Profile

	sqlScript := `SELECT id, account_id, name, address 
				FROM hydroponic_system.profiles 
				WHERE deleted_at IS NULL`

	res := r.db.Raw(sqlScript).Scan(&profiles)

	if res.Error != nil {
		return nil, res.Error
	}

	return profiles, nil

}

func (r profileRepository) CheckCreatedProfileByAccountId(inputModel *model.Profile) (*model.Profile, error) {

	sqlScript := `SELECT id, account_id, name, address 
				FROM hydroponic_system.profiles 
				WHERE account_id = ? AND deleted_at IS NOT NULL`

	res := r.db.Raw(sqlScript, inputModel.AccountId).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidAccountId
	}
	return inputModel, nil

}

func (r profileRepository) GetProfileById(inputModel *model.Profile) (*model.Profile, error) {

	sqlScript := `SELECT id, account_id, name, address 
				FROM hydroponic_system.profiles 
				WHERE id = ?`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidProfileID
	}
	return inputModel, nil

}

func (r profileRepository) UpdateProfile(inputModel *model.Profile) (*model.Profile, error) {

	sqlScript := `UPDATE hydroponic_system.profiles 
				SET updated_at = ?, name = ?, address = ?  
				WHERE id = ? 
				RETURNING id, account_id, name, address `

	res := r.db.Raw(sqlScript, time.Now(), inputModel.Name, inputModel.Address, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidProfileID
	}
	return inputModel, nil
}

func (r profileRepository) DeleteProfile(inputModel *model.Profile) (*model.Profile, error) {

	sqlScript := `UPDATE hydroponic_system.profiles 
				SET deleted_at = ? 
				WHERE id = ? 
				RETURNING id, account_id, name, address`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidProfileID
	}
	return inputModel, nil
}

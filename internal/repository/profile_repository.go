package repository

import (
	"time"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
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
func (r *profileRepository) CreateProfile(inputModel *model.Profile) (*model.Profile, error) {
	logger.Info("profileRepository", "Creating new profile", "account_id", inputModel.AccountId)

	sqlScript := `INSERT INTO hydroponic_system.profiles (account_id, name, address, created_at) 
				  VALUES (?, ?, ?, ?) 
				  RETURNING id, account_id, name, address;`

	res := r.db.Raw(sqlScript, inputModel.AccountId, inputModel.Name, inputModel.Address, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Failed to create profile", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("profileRepository", "Profile created successfully", "profile_id", inputModel.ID)
	return inputModel, nil
}

func (r *profileRepository) GetProfiles() ([]*model.Profile, error) {
	logger.Info("profileRepository", "Fetching all profiles")

	var profiles []*model.Profile

	sqlScript := `SELECT id, account_id, name, address 
				  FROM hydroponic_system.profiles 
				  WHERE deleted_at IS NULL`

	res := r.db.Raw(sqlScript).Scan(&profiles)

	if res.Error != nil {
		logger.Error("profileRepository", "Failed to fetch profiles", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("profileRepository", "Profiles fetched successfully", "count", len(profiles))
	return profiles, nil
}

func (r *profileRepository) CheckCreatedProfileByAccountId(inputModel *model.Profile) (*model.Profile, error) {
	logger.Info("profileRepository", "Checking if profile exists for account", "account_id", inputModel.AccountId)

	sqlScript := `SELECT id, account_id, name, address 
				  FROM hydroponic_system.profiles 
				  WHERE account_id = ? AND deleted_at IS NULL`

	res := r.db.Raw(sqlScript, inputModel.AccountId).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Error checking profile by account ID", "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("profileRepository", "No profile found for account", "account_id", inputModel.AccountId)
		return nil, errs.InvalidAccountId
	}

	logger.Info("profileRepository", "Profile found for account", "account_id", inputModel.AccountId)
	return inputModel, nil
}

func (r *profileRepository) GetProfileById(inputModel *model.Profile) (*model.Profile, error) {
	logger.Info("profileRepository", "Fetching profile by ID", "profile_id", inputModel.ID)

	sqlScript := `SELECT id, account_id, name, address 
				  FROM hydroponic_system.profiles 
				  WHERE id = ? AND deleted_at IS NULL`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Error fetching profile by ID", "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("profileRepository", "Profile not found", "profile_id", inputModel.ID)
		return nil, errs.InvalidProfileID
	}

	logger.Info("profileRepository", "Profile fetched successfully", "profile_id", inputModel.ID)
	return inputModel, nil
}

func (r *profileRepository) UpdateProfile(inputModel *model.Profile) (*model.Profile, error) {
	logger.Info("profileRepository", "Updating profile", "profile_id", inputModel.ID)

	sqlScript := `UPDATE hydroponic_system.profiles 
				  SET updated_at = ?, name = ?, address = ?  
				  WHERE id = ? AND deleted_at IS NULL
				  RETURNING id, account_id, name, address`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.Name, inputModel.Address, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Failed to update profile", "profile_id", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("profileRepository", "Profile not found for update", "profile_id", inputModel.ID)
		return nil, errs.InvalidProfileID
	}

	logger.Info("profileRepository", "Profile updated successfully", "profile_id", inputModel.ID)
	return inputModel, nil
}

func (r *profileRepository) DeleteProfile(inputModel *model.Profile) (*model.Profile, error) {
	logger.Info("profileRepository", "Deleting profile", "profile_id", inputModel.ID)

	sqlScript := `UPDATE hydroponic_system.profiles 
				  SET deleted_at = ? 
				  WHERE id = ? AND deleted_at IS NULL
				  RETURNING id, account_id, name, address`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Failed to delete profile", "profile_id", inputModel.ID, "error", res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("profileRepository", "Profile not found for deletion", "profile_id", inputModel.ID)
		return nil, errs.InvalidProfileID
	}

	logger.Info("profileRepository", "Profile deleted successfully", "profile_id", inputModel.ID)
	return inputModel, nil
}

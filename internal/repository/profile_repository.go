package repository

import (
	"strconv"
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
	logger.Info("profileRepository", "Creating new profile", map[string]string{
		"account_id": inputModel.AccountId.String(),
	})

	sqlScript := `INSERT INTO hydroponic_system.profiles (account_id, name, address, created_at) 
				  VALUES (?, ?, ?, ?) 
				  RETURNING id, account_id, name, address;`

	res := r.db.Raw(sqlScript, inputModel.AccountId, inputModel.Name, inputModel.Address, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Failed to create profile", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("profileRepository", "Profile created successfully", map[string]string{
		"profile_id": inputModel.ID.String(),
	})
	return inputModel, nil
}

func (r *profileRepository) GetProfiles() ([]*model.Profile, error) {
	logger.Info("profileRepository", "Fetching all profiles", nil)

	var profiles []*model.Profile

	sqlScript := `SELECT id, account_id, name, address 
				  FROM hydroponic_system.profiles 
				  WHERE deleted_at IS NULL`

	res := r.db.Raw(sqlScript).Scan(&profiles)

	if res.Error != nil {
		logger.Error("profileRepository", "Failed to fetch profiles", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("profileRepository", "Profiles fetched successfully", map[string]string{
		"count": strconv.Itoa(len(profiles)),
	})
	return profiles, nil
}

func (r *profileRepository) CheckCreatedProfileByAccountId(inputModel *model.Profile) (*model.Profile, error) {
	logger.Info("profileRepository", "Checking if profile exists for account", map[string]string{
		"account_id": inputModel.AccountId.String(),
	})

	sqlScript := `SELECT id, account_id, name, address 
				  FROM hydroponic_system.profiles 
				  WHERE account_id = ? AND deleted_at IS NULL`

	res := r.db.Raw(sqlScript, inputModel.AccountId).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Error checking profile by account ID", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("profileRepository", "No profile found for account", map[string]string{
			"account_id": inputModel.AccountId.String(),
		})
		return nil, errs.InvalidAccountId
	}

	logger.Info("profileRepository", "Profile found for account", map[string]string{
		"account_id": inputModel.AccountId.String(),
	})
	return inputModel, nil
}

func (r *profileRepository) GetProfileById(inputModel *model.Profile) (*model.Profile, error) {
	logger.Info("profileRepository", "Fetching profile by ID", map[string]string{
		"profile_id": inputModel.ID.String(),
	})

	sqlScript := `SELECT id, account_id, name, address 
				  FROM hydroponic_system.profiles 
				  WHERE id = ? AND deleted_at IS NULL`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Error fetching profile by ID", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("profileRepository", "Profile not found", map[string]string{
			"profile_id": inputModel.ID.String(),
		})
		return nil, errs.InvalidProfileID
	}

	logger.Info("profileRepository", "Profile fetched successfully", map[string]string{
		"profile_id": inputModel.ID.String(),
	})
	return inputModel, nil
}

func (r *profileRepository) UpdateProfile(inputModel *model.Profile) (*model.Profile, error) {
	logger.Info("profileRepository", "Updating profile", map[string]string{
		"profile_id": inputModel.ID.String(),
	})

	sqlScript := `UPDATE hydroponic_system.profiles 
				  SET updated_at = ?, name = ?, address = ?  
				  WHERE id = ? AND deleted_at IS NULL
				  RETURNING id, account_id, name, address`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.Name, inputModel.Address, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Failed to update profile", map[string]string{
			"profile_id": inputModel.ID.String(),
			"error":      res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("profileRepository", "Profile not found for update", map[string]string{
			"profile_id": inputModel.ID.String(),
		})
		return nil, errs.InvalidProfileID
	}

	logger.Info("profileRepository", "Profile updated successfully", map[string]string{
		"profile_id": inputModel.ID.String(),
	})
	return inputModel, nil
}

func (r *profileRepository) DeleteProfile(inputModel *model.Profile) (*model.Profile, error) {
	logger.Info("profileRepository", "Deleting profile", map[string]string{
		"profile_id": inputModel.ID.String(),
	})

	sqlScript := `UPDATE hydroponic_system.profiles 
				  SET deleted_at = ? 
				  WHERE id = ? AND deleted_at IS NULL
				  RETURNING id, account_id, name, address`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("profileRepository", "Failed to delete profile", map[string]string{
			"profile_id": inputModel.ID.String(),
			"error":      res.Error.Error(),
		})
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		logger.Warn("profileRepository", "Profile not found for deletion", map[string]string{
			"profile_id": inputModel.ID.String(),
		})
		return nil, errs.InvalidProfileID
	}

	logger.Info("profileRepository", "Profile deleted successfully", map[string]string{
		"profile_id": inputModel.ID.String(),
	})
	return inputModel, nil
}

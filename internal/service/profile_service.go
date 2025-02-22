package service

import (
	"strconv"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"

	"github.com/google/uuid"
)

type ProfileService interface {
	CreateProfile(input *dto.CreateProfile) (*dto.ProfileResponse, error)
	GetProfiles() ([]*dto.ProfileResponse, error)
	GetProfileDetails(profileId *uuid.UUID) (*dto.ProfileResponse, error)
	UpdateProfile(profileId *uuid.UUID, profileData *dto.UpdateProfile) (*dto.ProfileResponse, error)
	DeleteProfile(profileId *uuid.UUID) (*dto.ProfileResponse, error)
}

type profileService struct {
	profileRepo repository.ProfileRepository
	accountRepo repository.AccountRepository
}

type ProfileServiceConfig struct {
	ProfileRepo repository.ProfileRepository
	AccountRepo repository.AccountRepository
}

func NewProfileService(config ProfileServiceConfig) ProfileService {
	return &profileService{
		profileRepo: config.ProfileRepo,
		accountRepo: config.AccountRepo,
	}
}

func (s *profileService) CreateProfile(input *dto.CreateProfile) (*dto.ProfileResponse, error) {
	logger.Info("profileService", "Creating profile", map[string]string{
		"accountId": input.AccountID.String(),
	})

	user, err := s.accountRepo.GetUserById(input.AccountID)
	if err != nil || user == nil {
		logger.Error("profileService", "Invalid Account ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidAccountId
	}

	checkedProfile, err := s.profileRepo.CheckCreatedProfileByAccountId(&model.Profile{
		AccountId: input.AccountID,
	})
	if err != nil || checkedProfile == nil {
		logger.Error("profileService", "Profile already exists", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ProfileAlreadyCreated
	}

	createdProfile, err := s.profileRepo.CreateProfile(&model.Profile{
		AccountId: input.AccountID,
		Name:      input.Name,
		Address:   input.Address,
	})
	if err != nil {
		logger.Error("profileService", "Error creating profile", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorOnCreatingNewProfile
	}

	respBody := &dto.ProfileResponse{
		ID:      createdProfile.ID,
		Name:    createdProfile.Name,
		Address: createdProfile.Address,
	}

	logger.Info("profileService", "Profile created successfully", map[string]string{
		"profileId": respBody.ID.String(),
	})
	return respBody, err
}

func (s *profileService) GetProfiles() ([]*dto.ProfileResponse, error) {
	logger.Info("profileService", "Fetching all profiles", nil)

	var profilesRes []*dto.ProfileResponse

	res, err := s.profileRepo.GetProfiles()
	if err != nil {
		logger.Error("profileService", "Error fetching profiles", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		profilesRes = append(profilesRes, &dto.ProfileResponse{
			ID:      res[i].ID,
			Name:    res[i].Name,
			Address: res[i].Address,
		})
	}

	logger.Info("profileService", "Successfully fetched profiles", map[string]string{
		"count": strconv.Itoa(len(profilesRes)),
	})
	return profilesRes, err
}

func (s *profileService) GetProfileDetails(profileId *uuid.UUID) (*dto.ProfileResponse, error) {
	logger.Info("profileService", "Fetching profile details", map[string]string{
		"profileId": profileId.String(),
	})

	res, err := s.profileRepo.GetProfileById(&model.Profile{ID: *profileId})
	if err != nil {
		logger.Error("profileService", "Error fetching profile details", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("profileService", "Profile details fetched successfully", map[string]string{
		"profileId": profileId.String(),
	})
	return &dto.ProfileResponse{
		ID:      res.ID,
		Name:    res.Name,
		Address: res.Address,
	}, err
}

func (s *profileService) UpdateProfile(profileId *uuid.UUID, profileData *dto.UpdateProfile) (*dto.ProfileResponse, error) {
	logger.Info("profileService", "Updating profile", map[string]string{
		"profileId": profileId.String(),
	})

	res, err := s.profileRepo.UpdateProfile(&model.Profile{ID: *profileId, Name: profileData.Name, Address: profileData.Address})
	if err != nil {
		logger.Error("profileService", "Error updating profile", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("profileService", "Profile updated successfully", map[string]string{
		"profileId": profileId.String(),
	})
	return &dto.ProfileResponse{
		ID:      res.ID,
		Name:    res.Name,
		Address: res.Address,
	}, err
}

func (s *profileService) DeleteProfile(profileId *uuid.UUID) (*dto.ProfileResponse, error) {
	logger.Info("profileService", "Deleting profile", map[string]string{
		"profileId": profileId.String(),
	})

	res, err := s.profileRepo.DeleteProfile(&model.Profile{ID: *profileId})
	if err != nil {
		logger.Error("profileService", "Error deleting profile", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("profileService", "Profile deleted successfully", map[string]string{
		"profileId": profileId.String(),
	})
	return &dto.ProfileResponse{
		ID: res.ID,
	}, err
}

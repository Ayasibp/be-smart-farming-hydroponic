package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/google/uuid"
)

type ProfileService interface {
	CreateProfile(input *dto.CreateProfile) (*dto.ProfileResponse, error)
	GetProfileDetails(profileId *uuid.UUID) (*dto.ProfileResponse, error)
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

func (ps profileService) CreateProfile(input *dto.CreateProfile) (*dto.ProfileResponse, error) {

	user, err := ps.accountRepo.GetUserById(input.AccountID)
	if err != nil || user == nil {
		return nil, errs.InvalidAccountId
	}

	createdProfile, err := ps.profileRepo.CreateProfile(&model.Profile{
		AccountId: input.AccountID,
		Name:      input.Name,
		Address:   input.Address,
	})
	if err != nil {
		return nil, errs.ErrorOnCreatingNewProfile
	}

	respBody := &dto.ProfileResponse{
		ID:      createdProfile.ID,
		Name:    createdProfile.Name,
		Address: createdProfile.Address,
	}

	return respBody, err
}

func (ps profileService) GetProfileDetails(profileId *uuid.UUID) (*dto.ProfileResponse, error) {

	res, err := ps.profileRepo.GetProfileById(&model.Profile{ID: *profileId})
	if err != nil {
		return nil, err
	}

	return &dto.ProfileResponse{
		ID:      res.ID,
		Name:    res.Name,
		Address: res.Address,
	}, err
}

func (ps profileService) DeleteProfile(profileId *uuid.UUID) (*dto.ProfileResponse, error) {

	res, err := ps.profileRepo.DeleteProfile(&model.Profile{ID: *profileId})
	if err != nil {
		return nil, err
	}

	return &dto.ProfileResponse{
		ID: res.ID,
	}, err
}

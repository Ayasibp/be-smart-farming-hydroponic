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

func (s profileService) CreateProfile(input *dto.CreateProfile) (*dto.ProfileResponse, error) {

	user, err := s.accountRepo.GetUserById(input.AccountID)
	if err != nil || user == nil {
		return nil, errs.InvalidAccountId
	}

	checkedProfile, err := s.profileRepo.CheckCreatedProfileByAccountId(&model.Profile{
		AccountId: input.AccountID,
	})
	if err != nil || checkedProfile == nil {
		return nil, errs.ProfileAlreadyCreated
	}

	createdProfile, err := s.profileRepo.CreateProfile(&model.Profile{
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

func (s profileService) GetProfiles() ([]*dto.ProfileResponse, error) {

	var profilesRes []*dto.ProfileResponse

	res, err := s.profileRepo.GetProfiles()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		profilesRes = append(profilesRes, &dto.ProfileResponse{
			ID:      res[i].ID,
			Name:    res[i].Name,
			Address: res[i].Address,
		})
	}

	return profilesRes, err
}

func (s profileService) GetProfileDetails(profileId *uuid.UUID) (*dto.ProfileResponse, error) {

	res, err := s.profileRepo.GetProfileById(&model.Profile{ID: *profileId})
	if err != nil {
		return nil, err
	}

	return &dto.ProfileResponse{
		ID:      res.ID,
		Name:    res.Name,
		Address: res.Address,
	}, err
}
func (s profileService) UpdateProfile(profileId *uuid.UUID, profileData *dto.UpdateProfile) (*dto.ProfileResponse, error) {

	res, err := s.profileRepo.UpdateProfile(&model.Profile{ID: *profileId, Name: profileData.Name, Address: profileData.Address})
	if err != nil {
		return nil, err
	}

	return &dto.ProfileResponse{
		ID:      res.ID,
		Name:    res.Name,
		Address: res.Address,
	}, err
}

func (s profileService) DeleteProfile(profileId *uuid.UUID) (*dto.ProfileResponse, error) {

	res, err := s.profileRepo.DeleteProfile(&model.Profile{ID: *profileId})
	if err != nil {
		return nil, err
	}

	return &dto.ProfileResponse{
		ID: res.ID,
	}, err
}

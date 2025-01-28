package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type ProfileService interface {
	CreateProfile(input dto.CreateProfile) (*dto.ProfileResponse, error)
}

type profileService struct {
	profileRepo repository.ProfileRepository
}

type ProfileServiceConfig struct {
	ProfileRepo repository.ProfileRepository
}

func NewProfileService(config ProfileServiceConfig) ProfileService {
	return &profileService{
		profileRepo: config.ProfileRepo,
	}
}

func (ps profileService) CreateProfile(input dto.CreateProfile) (*dto.ProfileResponse, error) {
	
	createdProfile, err := ps.profileRepo.CreateProfile(&dto.CreateProfile{
		AccountID : input.AccountID,
		Name: input.Name,
		Address: input.Address,
	})

	respBody := &dto.ProfileResponse{
		ID:   createdProfile.ID,
		Name: createdProfile.Name,
		Address:     createdProfile.Address,
	}

	return respBody, err
}
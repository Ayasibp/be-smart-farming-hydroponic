package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type ProfileService interface {
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

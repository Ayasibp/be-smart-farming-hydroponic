package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/google/uuid"
)

type FarmService interface {
	CreateFarm(input *dto.CreateFarm) (*dto.FarmResponse, error)
	GetFarms() ([]*dto.FarmResponse, error)
	GetFarmDetails(farmId *uuid.UUID) (*dto.FarmResponse, error)
}

type farmService struct {
	farmRepo    repository.FarmRepository
	profileRepo repository.ProfileRepository
}

type FarmServiceConfig struct {
	FarmRepo    repository.FarmRepository
	ProfileRepo repository.ProfileRepository
}

func NewFarmService(config FarmServiceConfig) FarmService {
	return &farmService{
		farmRepo:    config.FarmRepo,
		profileRepo: config.ProfileRepo,
	}
}

func (s farmService) CreateFarm(input *dto.CreateFarm) (*dto.FarmResponse, error) {

	profile, err := s.profileRepo.GetProfileById(&model.Profile{ID: input.ProfileID})
	if err != nil || profile == nil {
		return nil, errs.InvalidProfileID
	}

	createdFarm, err := s.farmRepo.CreateFarm(&model.Farm{
		ProfileId: input.ProfileID,
		Name:      input.Name,
		Address:   input.Address,
	})
	if err != nil {
		return nil, errs.ErrorOnCreatingNewFarm
	}

	respBody := &dto.FarmResponse{
		ID:      createdFarm.ID,
		Name:    createdFarm.Name,
		Address: createdFarm.Address,
	}

	return respBody, err
}

func (s farmService) GetFarms() ([]*dto.FarmResponse, error) {

	var farmResponse []*dto.FarmResponse

	res, err := s.farmRepo.GetFarms()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		farmResponse = append(farmResponse, &dto.FarmResponse{
			ID:      res[i].ID,
			Name:    res[i].Name,
			Address: res[i].Address,
		})
	}

	return farmResponse, err
}

func (s farmService) GetFarmDetails(farmId *uuid.UUID) (*dto.FarmResponse, error) {

	res, err := s.farmRepo.GetFarmById(&model.Farm{ID: *farmId})
	if err != nil {
		return nil, err
	}

	return &dto.FarmResponse{
		ID:      res.ID,
		Name:    res.Name,
		Address: res.Address,
	}, err
}

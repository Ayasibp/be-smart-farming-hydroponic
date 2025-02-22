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

type FarmService interface {
	CreateFarm(input *dto.CreateFarm) (*dto.FarmResponse, error)
	GetFarms() ([]*dto.FarmResponse, error)
	GetFarmDetails(farmId *uuid.UUID) (*dto.FarmResponse, error)
	UpdateFarm(farmId *uuid.UUID, farmData *dto.UpdateFarm) (*dto.FarmResponse, error)
	DeleteFarm(farmId *uuid.UUID) (*dto.FarmResponse, error)
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

func (s *farmService) CreateFarm(input *dto.CreateFarm) (*dto.FarmResponse, error) {
	logger.Info("farmService", "Creating new farm", map[string]string{
		"profile_id": input.ProfileID.String(),
		"name":       input.Name,
	})

	profile, err := s.profileRepo.GetProfileById(&model.Profile{ID: input.ProfileID})
	if err != nil || profile == nil {
		logger.Error("farmService", "Invalid profile ID", map[string]string{
			"profile_id": input.ProfileID.String(),
			"error":      err.Error(),
		})
		return nil, errs.InvalidProfileID
	}

	createdFarm, err := s.farmRepo.CreateFarm(&model.Farm{
		ProfileId: input.ProfileID,
		Name:      input.Name,
		Address:   input.Address,
	})
	if err != nil {
		logger.Error("farmService", "Error creating farm", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorOnCreatingNewFarm
	}

	logger.Info("farmService", "Farm created successfully", map[string]string{
		"farm_id": createdFarm.ID.String(),
		"name":    createdFarm.Name,
	})
	return &dto.FarmResponse{
		ID:      createdFarm.ID,
		Name:    createdFarm.Name,
		Address: createdFarm.Address,
	}, nil
}

func (s *farmService) GetFarms() ([]*dto.FarmResponse, error) {
	logger.Info("farmService", "Fetching all farms", nil)

	res, err := s.farmRepo.GetFarms()
	if err != nil {
		logger.Error("farmService", "Failed to fetch farms", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	var farmResponse []*dto.FarmResponse
	for _, farm := range res {
		farmResponse = append(farmResponse, &dto.FarmResponse{
			ID:      farm.ID,
			Name:    farm.Name,
			Address: farm.Address,
		})
	}

	logger.Info("farmService", "Fetched farms successfully", map[string]string{
		"count": strconv.Itoa(len(farmResponse)),
	})
	return farmResponse, nil
}

func (s *farmService) GetFarmDetails(farmId *uuid.UUID) (*dto.FarmResponse, error) {
	logger.Info("farmService", "Fetching farm details", map[string]string{
		"farm_id": farmId.String(),
	})

	res, err := s.farmRepo.GetFarmById(&model.Farm{ID: *farmId})
	if err != nil {
		logger.Error("farmService", "Failed to fetch farm details", map[string]string{
			"farm_id": farmId.String(),
			"error":   err.Error(),
		})
		return nil, err
	}

	logger.Info("farmService", "Fetched farm details successfully", map[string]string{
		"farm_id": res.ID.String(),
		"name":    res.Name,
	})
	return &dto.FarmResponse{
		ID:      res.ID,
		Name:    res.Name,
		Address: res.Address,
	}, nil
}

func (s *farmService) UpdateFarm(farmId *uuid.UUID, farmData *dto.UpdateFarm) (*dto.FarmResponse, error) {
	logger.Info("farmService", "Updating farm", map[string]string{
		"farm_id":  farmId.String(),
		"new_name": farmData.Name,
	})

	res, err := s.farmRepo.UpdateFarm(&model.Farm{ID: *farmId, Name: farmData.Name, Address: farmData.Address})
	if err != nil {
		logger.Error("farmService", "Failed to update farm", map[string]string{
			"farm_id": farmId.String(),
			"error":   err.Error(),
		})
		return nil, err
	}

	logger.Info("farmService", "Farm updated successfully", map[string]string{
		"farm_id":  res.ID.String(),
		"new_name": res.Name,
	})
	return &dto.FarmResponse{
		ID:      res.ID,
		Name:    res.Name,
		Address: res.Address,
	}, nil
}

func (s *farmService) DeleteFarm(farmId *uuid.UUID) (*dto.FarmResponse, error) {
	logger.Info("farmService", "Deleting farm", map[string]string{
		"farm_id": farmId.String(),
	})

	res, err := s.farmRepo.DeleteFarm(&model.Farm{ID: *farmId})
	if err != nil {
		logger.Error("farmService", "Failed to delete farm", map[string]string{
			"farm_id": farmId.String(),
			"error":   err.Error(),
		})
		return nil, err
	}

	logger.Info("farmService", "Farm deleted successfully", map[string]string{
		"farm_id": res.ID.String(),
	})
	return &dto.FarmResponse{
		ID: res.ID,
	}, nil
}

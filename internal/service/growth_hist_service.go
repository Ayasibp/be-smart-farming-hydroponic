package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type GrowthHistService interface {
	CreateGrowthHist(input *dto.GrowthHist) (*dto.GrowthHistResponse, error)
}

type growthHistService struct {
	growthHistRepo repository.GrowthHistRepository
	farmRepo       repository.FarmRepository
	systemUnitRepo repository.SystemUnitRepository
}

type GrowthHistServiceConfig struct {
	GrowthHistRepo repository.GrowthHistRepository
	FarmRepo       repository.FarmRepository
	SystemUnitRepo repository.SystemUnitRepository
}

func NewGrowthHistService(config GrowthHistServiceConfig) GrowthHistService {
	return &growthHistService{
		growthHistRepo: config.GrowthHistRepo,
		farmRepo:       config.FarmRepo,
		systemUnitRepo: config.SystemUnitRepo,
	}
}

func (s growthHistService) CreateGrowthHist(input *dto.GrowthHist) (*dto.GrowthHistResponse, error) {

	farm, err := s.farmRepo.GetFarmById(&model.Farm{
		ID: input.FarmId,
	})
	if err != nil || farm == nil {
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{
		ID: input.SystemId,
	})
	if err != nil || systemUnit == nil {
		return nil, errs.InvalidSystemUnitID
	}

	growthHist, err := s.growthHistRepo.CreateGrowthHistory(&model.GrowthHist{
		FarmId:   input.FarmId,
		SystemId: input.SystemId,
		Ppm:      input.Ppm,
		Ph:       input.Ph,
	})
	if err != nil {
		return nil, errs.ErrorOnCreatingNewProfile
	}

	respBody := &dto.GrowthHistResponse{
		ID:       growthHist.ID,
		FarmId:   growthHist.FarmId,
		SystemId: growthHist.SystemId,
		Ppm:      growthHist.Ppm,
		Ph:       growthHist.Ph,
	}

	return respBody, err
}

package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type SystemUnitService interface {
	CreateSystemUnit(input *dto.CreateSystemUnit) (*dto.SystemUnitResponse, error)
}

type systemUnitService struct {
	systemUnitRepo    repository.SystemUnitRepository
	farmRepo repository.FarmRepository
}

type SystemUnitServiceConfig struct {
	SystemUnitRepo    repository.SystemUnitRepository
	FarmRepo repository.FarmRepository
}

func NewSystemUnitService(config SystemUnitServiceConfig) SystemUnitService {
	return &systemUnitService{
		systemUnitRepo:    config.SystemUnitRepo,
		farmRepo: config.FarmRepo,
		
	}
}

func (s systemUnitService) CreateSystemUnit(input *dto.CreateSystemUnit) (*dto.SystemUnitResponse, error) {

	farm, err:= s.farmRepo.GetFarmById(&model.Farm{ID: input.FarmID})
	if err != nil || farm == nil {
		return nil, errs.InvalidFarmID
	}

	createdSystemUnit, err := s.systemUnitRepo.CreateSystemUnit(&model.SystemUnit{
		FarmId: input.FarmID,
		TankVolume:      input.TankVolume,
		TankAVolume:   input.TankAVolume,
		TankBVolume: input.TankBVolume,
	})
	if err != nil {
		return nil, errs.ErrorOnCreatingNewSystemUnit
	}

	respBody := &dto.SystemUnitResponse{
		ID:      createdSystemUnit.ID,
		TankVolume:    createdSystemUnit.TankVolume,
		TankAVolume: createdSystemUnit.TankAVolume,
		TankBVolume: createdSystemUnit.TankBVolume,
	}

	return respBody, err
}
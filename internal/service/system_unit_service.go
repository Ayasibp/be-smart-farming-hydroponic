package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/google/uuid"
)

type SystemUnitService interface {
	CreateSystemUnit(input *dto.CreateSystemUnit) (*dto.CreateSystemUnitResponse, error)
	DeleteSystemUnitById(unitId *uuid.UUID) (*dto.CreateSystemUnitResponse, error)
}

type systemUnitService struct {
	systemUnitRepo repository.SystemUnitRepository
	farmRepo       repository.FarmRepository
	unitKeyRepo    repository.UnitIdRepository
}

type SystemUnitServiceConfig struct {
	SystemUnitRepo repository.SystemUnitRepository
	FarmRepo       repository.FarmRepository
	UnitKeyRepo    repository.UnitIdRepository
}

func NewSystemUnitService(config SystemUnitServiceConfig) SystemUnitService {
	return &systemUnitService{
		systemUnitRepo: config.SystemUnitRepo,
		farmRepo:       config.FarmRepo,
		unitKeyRepo:    config.UnitKeyRepo,
	}
}

func (s systemUnitService) CreateSystemUnit(input *dto.CreateSystemUnit) (*dto.CreateSystemUnitResponse, error) {

	farm, err := s.farmRepo.GetFarmById(&model.Farm{ID: input.FarmID})
	if err != nil || farm == nil {
		return nil, errs.InvalidFarmID
	}
	unitKey, err := s.unitKeyRepo.GetUnitIdById(&model.UnitId{ID: input.UnitKey})
	if err != nil || unitKey == nil {
		return nil, errs.InvalidUnitKey
	}

	createdSystemUnit, err := s.systemUnitRepo.CreateSystemUnit(&model.SystemUnit{
		FarmId:      input.FarmID,
		UnitKey:     input.UnitKey,
		TankVolume:  input.TankVolume,
		TankAVolume: input.TankAVolume,
		TankBVolume: input.TankBVolume,
	})
	if err != nil {
		return nil, errs.ErrorOnCreatingNewSystemUnit
	}

	respBody := &dto.CreateSystemUnitResponse{
		ID:          createdSystemUnit.ID,
		TankVolume:  createdSystemUnit.TankVolume,
		TankAVolume: createdSystemUnit.TankAVolume,
		TankBVolume: createdSystemUnit.TankBVolume,
	}

	return respBody, err
}

func (s systemUnitService) DeleteSystemUnitById(unitId *uuid.UUID) (*dto.CreateSystemUnitResponse, error) {

	res, err := s.systemUnitRepo.DeleteSystemUnitById(&model.SystemUnit{ID: *unitId})
	if err != nil {
		return nil, err
	}

	return &dto.CreateSystemUnitResponse{
		ID: res.ID,
	}, err
}
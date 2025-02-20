package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/google/uuid"
)

type SystemUnitService interface {
	CreateSystemUnit(input *dto.CreateSystemUnit) (*dto.CreateSystemUnitResponse, error)
	GetSystemUnits(farm_ids *dto.SystemUnitFilter) ([]*dto.SystemUnitResponse, error)
	UpdateSystemUnit(systemUnitId *uuid.UUID, systemUnitData *dto.CreateSystemUnit) (*dto.SystemUnitResponse, error)
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

func (s *systemUnitService) CreateSystemUnit(input *dto.CreateSystemUnit) (*dto.CreateSystemUnitResponse, error) {
	logger.Info("systemUnitService", "Creating a new system unit", "farm_id", input.FarmID, "unit_key", input.UnitKey)

	farm, err := s.farmRepo.GetFarmById(&model.Farm{ID: input.FarmID})
	if err != nil || farm == nil {
		logger.Error("systemUnitService", "Invalid farm ID", "error", err)
		return nil, errs.InvalidFarmID
	}

	unitKey, err := s.unitKeyRepo.GetUnitIdById(&model.UnitId{ID: input.UnitKey})
	if err != nil || unitKey == nil {
		logger.Error("systemUnitService", "Invalid unit key", "error", err)
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
		logger.Error("systemUnitService", "Error creating new system unit", "error", err)
		return nil, errs.ErrorOnCreatingNewSystemUnit
	}

	logger.Info("systemUnitService", "System unit created successfully", "unit_id", createdSystemUnit.ID)

	return &dto.CreateSystemUnitResponse{
		ID:          createdSystemUnit.ID,
		TankVolume:  createdSystemUnit.TankVolume,
		TankAVolume: createdSystemUnit.TankAVolume,
		TankBVolume: createdSystemUnit.TankBVolume,
	}, nil
}

func (s *systemUnitService) GetSystemUnits(farm_ids *dto.SystemUnitFilter) ([]*dto.SystemUnitResponse, error) {
	logger.Info("systemUnitService", "Fetching system units", "farm_ids", farm_ids)

	var systemUnitRes []*dto.SystemUnitResponse
	var id string

	if farm_ids == nil {
		id = ""
	} else {
		id = "AND su.farm_id IN (" + farm_ids.FarmIds + ")"
	}

	res, err := s.systemUnitRepo.GetSystemUnits(&id)
	if err != nil {
		logger.Error("systemUnitService", "Error fetching system units", "error", err)
		return nil, err
	}

	for _, resIdx := range res {
		systemUnitRes = append(systemUnitRes, &dto.SystemUnitResponse{
			ID:          resIdx.ID,
			UnitKey:     resIdx.UnitKey,
			FarmID:      resIdx.FarmId,
			FarmName:    resIdx.FarmName,
			TankVolume:  resIdx.TankVolume,
			TankAVolume: resIdx.TankAVolume,
			TankBVolume: resIdx.TankBVolume,
		})
	}

	logger.Info("systemUnitService", "Successfully fetched system units", "count", len(systemUnitRes))
	return systemUnitRes, nil
}

func (s *systemUnitService) UpdateSystemUnit(systemUnitId *uuid.UUID, systemUnitData *dto.CreateSystemUnit) (*dto.SystemUnitResponse, error) {
	logger.Info("systemUnitService", "Updating system unit", "unit_id", systemUnitId)

	res, err := s.systemUnitRepo.UpdateSystemUnit(&model.SystemUnit{
		ID:          *systemUnitId,
		FarmId:      systemUnitData.FarmID,
		UnitKey:     systemUnitData.UnitKey,
		TankVolume:  systemUnitData.TankVolume,
		TankAVolume: systemUnitData.TankAVolume,
		TankBVolume: systemUnitData.TankBVolume,
	})
	if err != nil {
		logger.Error("systemUnitService", "Error updating system unit", "error", err)
		return nil, err
	}

	logger.Info("systemUnitService", "System unit updated successfully", "unit_id", res.ID)
	return &dto.SystemUnitResponse{
		ID:          res.ID,
		FarmID:      res.FarmId,
		UnitKey:     res.UnitKey,
		TankVolume:  res.TankVolume,
		TankAVolume: res.TankAVolume,
		TankBVolume: res.TankBVolume,
	}, nil
}

func (s *systemUnitService) DeleteSystemUnitById(unitId *uuid.UUID) (*dto.CreateSystemUnitResponse, error) {
	logger.Info("systemUnitService", "Deleting system unit", "unit_id", unitId)

	res, err := s.systemUnitRepo.DeleteSystemUnitById(&model.SystemUnit{ID: *unitId})
	if err != nil {
		logger.Error("systemUnitService", "Error deleting system unit", "error", err)
		return nil, err
	}

	logger.Info("systemUnitService", "System unit deleted successfully", "unit_id", res.ID)
	return &dto.CreateSystemUnitResponse{
		ID: res.ID,
	}, nil
}

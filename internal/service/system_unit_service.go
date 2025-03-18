package service

import (
	"strconv"

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
	logger.Info("systemUnitService", "Creating a new system unit", map[string]string{
		"farm_id":  input.FarmID.String(),
		"unit_key": input.UnitKey.String(),
	})

	farm, err := s.farmRepo.GetFarmById(&model.Farm{ID: input.FarmID})
	if err != nil || farm == nil {
		logger.Error("systemUnitService", "Invalid farm ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidFarmID
	}

	unitKey, err := s.unitKeyRepo.GetUnitIdById(&model.UnitId{ID: input.UnitKey})
	if err != nil || unitKey == nil {
		logger.Error("systemUnitService", "Invalid unit key", map[string]string{
			"error": err.Error(),
		})
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
		logger.Error("systemUnitService", "Error creating new system unit", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorOnCreatingNewSystemUnit
	}

	logger.Info("systemUnitService", "System unit created successfully", map[string]string{
		"unit_id": createdSystemUnit.ID.String(),
	})

	return &dto.CreateSystemUnitResponse{
		ID:          createdSystemUnit.ID,
		TankVolume:  createdSystemUnit.TankVolume,
		TankAVolume: createdSystemUnit.TankAVolume,
		TankBVolume: createdSystemUnit.TankBVolume,
	}, nil
}

func (s *systemUnitService) GetSystemUnits(farm_ids *dto.SystemUnitFilter) ([]*dto.SystemUnitResponse, error) {
	logger.Info("systemUnitService", "Fetching system units", map[string]string{
		"farm_ids": farm_ids.FarmIds,
	})

	var systemUnitRes []*dto.SystemUnitResponse
	var id string

	if farm_ids == nil {
		id = ""
	} else {
		id = "AND su.farm_id IN (" + farm_ids.FarmIds + ")"
	}

	res, err := s.systemUnitRepo.GetSystemUnits(&id)
	if err != nil {
		logger.Error("systemUnitService", "Error fetching system units", map[string]string{
			"error": err.Error(),
		})
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

	logger.Info("systemUnitService", "Successfully fetched system units", map[string]string{
		"count": strconv.Itoa(len(systemUnitRes)),
	})
	return systemUnitRes, nil
}

func (s *systemUnitService) UpdateSystemUnit(systemUnitId *uuid.UUID, systemUnitData *dto.CreateSystemUnit) (*dto.SystemUnitResponse, error) {
	logger.Info("systemUnitService", "Updating system unit", map[string]string{
		"unit_id": systemUnitId.String(),
	})

	res, err := s.systemUnitRepo.UpdateSystemUnit(&model.SystemUnit{
		ID:          *systemUnitId,
		FarmId:      systemUnitData.FarmID,
		UnitKey:     systemUnitData.UnitKey,
		TankVolume:  systemUnitData.TankVolume,
		TankAVolume: systemUnitData.TankAVolume,
		TankBVolume: systemUnitData.TankBVolume,
	})
	if err != nil {
		logger.Error("systemUnitService", "Error updating system unit", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("systemUnitService", "System unit updated successfully", map[string]string{
		"unit_id": systemUnitId.String(),
	})
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
	logger.Info("systemUnitService", "Deleting system unit", map[string]string{
		"unit_id": unitId.String(),
	})

	res, err := s.systemUnitRepo.DeleteSystemUnitById(&model.SystemUnit{ID: *unitId})
	if err != nil {
		logger.Error("systemUnitService", "Error deleting system unit", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("systemUnitService", "System unit deleted successfully", map[string]string{
		"unit_id": unitId.String(),
	})
	return &dto.CreateSystemUnitResponse{
		ID: res.ID,
	}, nil
}

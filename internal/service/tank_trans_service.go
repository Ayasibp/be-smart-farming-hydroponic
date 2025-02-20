package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type TankTransService interface {
	CreateTankTrans(input *dto.TankTransaction) (*dto.TankTransactionResponse, error)
}

type tankTransService struct {
	tankTransRepo  repository.TankTransRepository
	farmRepo       repository.FarmRepository
	systemUnitRepo repository.SystemUnitRepository
}

type TankTransServiceConfig struct {
	TankTransRepo  repository.TankTransRepository
	FarmRepo       repository.FarmRepository
	SystemUnitRepo repository.SystemUnitRepository
}

func NewTankTransService(config TankTransServiceConfig) TankTransService {
	return &tankTransService{
		tankTransRepo:  config.TankTransRepo,
		farmRepo:       config.FarmRepo,
		systemUnitRepo: config.SystemUnitRepo,
	}
}

func (s *tankTransService) CreateTankTrans(input *dto.TankTransaction) (*dto.TankTransactionResponse, error) {
	logger.Info("tankTransService", "Creating tank transaction", "farm_id", input.FarmId, "system_id", input.SystemId)

	farm, err := s.farmRepo.GetFarmById(&model.Farm{
		ID: input.FarmId,
	})
	if err != nil || farm == nil {
		logger.Error("tankTransService", "Invalid farm ID", "farm_id", input.FarmId)
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{
		ID: input.SystemId,
	})
	if err != nil || systemUnit == nil {
		logger.Error("tankTransService", "Invalid system unit ID", "system_id", input.SystemId)
		return nil, errs.InvalidSystemUnitID
	}

	tankTrans, err := s.tankTransRepo.CreateTankTransaction(&model.TankTran{
		FarmId:      input.FarmId,
		SystemId:    input.SystemId,
		WaterVolume: input.WaterVolume,
		AVolume:     input.AVolume,
		BVolume:     input.BVolume,
	})
	if err != nil {
		logger.Error("tankTransService", "Error creating new tank transaction", "farm_id", input.FarmId, "system_id", input.SystemId)
		return nil, errs.ErrorOnCreatingNewTankTrans
	}

	logger.Info("tankTransService", "Successfully created tank transaction", "transaction_id", tankTrans.ID)

	respBody := &dto.TankTransactionResponse{
		ID:          tankTrans.ID,
		FarmId:      tankTrans.FarmId,
		SystemId:    tankTrans.SystemId,
		WaterVolume: tankTrans.WaterVolume,
		AVolume:     tankTrans.AVolume,
		BVolume:     tankTrans.BVolume,
	}

	return respBody, err
}

package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type TankTransService interface {
	CreateTankTrans(input *dto.TankTransaction) (*dto.TankTransactionResponse, error)
}

type tankTransService struct {
	tankTransRepo repository.TankTransRepository
	farmRepo       repository.FarmRepository
	systemUnitRepo repository.SystemUnitRepository
}

type TankTransServiceConfig struct {
	TankTransRepo repository.TankTransRepository
	FarmRepo       repository.FarmRepository
	SystemUnitRepo repository.SystemUnitRepository
}

func NewTankTransService(config TankTransServiceConfig) TankTransService {
	return &tankTransService{
		tankTransRepo: config.TankTransRepo,
		farmRepo:       config.FarmRepo,
		systemUnitRepo: config.SystemUnitRepo,
	}
}

func (s tankTransService) CreateTankTrans(input *dto.TankTransaction) (*dto.TankTransactionResponse, error) {

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

	tankTrans, err := s.tankTransRepo.CreateTankTransaction(&model.TankTran{
		FarmId:   input.FarmId,
		SystemId: input.SystemId,
		WaterVolume:      input.WaterVolume,
		AVolume:       input.AVolume,
		BVolume: input.BVolume,
	})
	if err != nil {
		return nil, errs.ErrorOnCreatingNewTankTrans
	}

	respBody := &dto.TankTransactionResponse{
		ID:       tankTrans.ID,
		FarmId:   tankTrans.FarmId,
		SystemId: tankTrans.SystemId,
		WaterVolume:      tankTrans.WaterVolume,
		AVolume:       tankTrans.AVolume,
		BVolume: tankTrans.BVolume,
	}

	return respBody, err
}
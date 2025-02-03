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
	GetSystemUnits(farm_ids *dto.SystemUnitFilter) ([]*dto.SystemUnitResponse, error)
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

func (s systemUnitService) GetSystemUnits(farm_ids *dto.SystemUnitFilter) ([]*dto.SystemUnitResponse, error) {

	var systemUnitRes []*dto.SystemUnitResponse
	var id string
	
	if farm_ids == nil{
		id = ""
	}else{
		id = "AND su.farm_id IN (" + farm_ids.FarmIds + ")"
	}
	
	res, err := s.systemUnitRepo.GetSystemUnits(&id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(res); i++ {
		resIdx := res[i]
		systemUnitRes = append(systemUnitRes, &dto.SystemUnitResponse{
			ID: resIdx.ID,
			UnitKey: resIdx.UnitKey,
			FarmID: resIdx.FarmId,
			FarmName: resIdx.FarmName,
			TankVolume: resIdx.TankVolume,
			TankAVolume: resIdx.TankAVolume,
			TankBVolume: resIdx.TankBVolume,
		})
	}

	return systemUnitRes, nil
}

func (s systemUnitService) UpdateSystemUnit(systemUnitId *uuid.UUID, systemUnitData *dto.CreateSystemUnit) (*dto.CreateSystemUnit, error) {

	res, err := s.systemUnitRepo.UpdateSystemUnit(&model.SystemUnit{ID: *systemUnitId, FarmId: systemUnitData.FarmID,UnitKey: systemUnitData.UnitKey,TankVolume: systemUnitData.TankVolume, TankAVolume: systemUnitData.TankAVolume, TankBVolume: systemUnitData.TankBVolume})
	if err != nil {
		return nil, err
	}

	return &dto.CreateSystemUnit{
		FarmID:      res.FarmId,
		UnitKey:    res.UnitKey,
		TankVolume: res.TankVolume,
		TankAVolume: res.TankAVolume,
		TankBVolume: res.TankBVolume,
	}, err
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
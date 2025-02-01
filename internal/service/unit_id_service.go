package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/google/uuid"
)

type UnitIdService interface {
	CreateUnitId() (*dto.UnitIdResponse, error)
	GetUnitIds() ([]*dto.UnitIdResponse, error)
	DeleteUnitIdbyId(unitId *uuid.UUID) (*dto.UnitIdResponse, error)
}

type unitIdService struct {
	unitIdRepo repository.UnitIdRepository
}

type UnitIdServiceConfig struct {
	UnitIdRepo repository.UnitIdRepository
}

func NewUnitIdService(config UnitIdServiceConfig) UnitIdService {
	return &unitIdService{
		unitIdRepo: config.UnitIdRepo,
	}
}

func (s unitIdService) CreateUnitId() (*dto.UnitIdResponse, error) {

	res, err := s.unitIdRepo.CreateUnitId()
	if err != nil {
		return nil, errs.ErrorCreatingAccount
	}

	return &dto.UnitIdResponse{
		ID: res.ID,
	}, err

}
func (s unitIdService) GetUnitIds() ([]*dto.UnitIdResponse, error) {

	var unitIdRes []*dto.UnitIdResponse

	res, err := s.unitIdRepo.GetUnitIds()
	if err != nil {
		return nil, errs.ErrorCreatingAccount
	}

	for i := 0; i < len(res); i++ {
		unitIdRes = append(unitIdRes, &dto.UnitIdResponse{
			ID: res[i].ID,
		})
	}

	return unitIdRes, err

}

func (s unitIdService) DeleteUnitIdbyId(unitId *uuid.UUID) (*dto.UnitIdResponse, error) {

	res, err := s.unitIdRepo.DeleteUnitIdById(&model.UnitId{ID: *unitId})
	if err != nil {
		return nil, err
	}

	return &dto.UnitIdResponse{
		ID: res.ID,
	}, err
}

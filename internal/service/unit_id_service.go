package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type UnitIdService interface {
	CreateUnitId() (*dto.UnitIdResponse, error)
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

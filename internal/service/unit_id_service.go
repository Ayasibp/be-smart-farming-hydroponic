package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type UnitIdService interface {
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

package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type FarmService interface {
}

type farmService struct {
	farmRepo repository.FarmRepository
}

type FarmServiceConfig struct {
	FarmRepo repository.FarmRepository
}

func NewFarmService(config FarmServiceConfig) FarmService {
	return &farmService{
		farmRepo: config.FarmRepo,
	}
}

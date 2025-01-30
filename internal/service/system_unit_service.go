package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type SystemUnitService interface {
	
}

type systemUnitService struct {
	systemUnitRepo    repository.SystemUnitRepository
}

type SystemUnitServiceConfig struct {
	SystemUnitRepo    repository.SystemUnitRepository
}

func NewSystemUnitService(config SystemUnitServiceConfig) SystemUnitService {
	return &systemUnitService{
		systemUnitRepo:    config.SystemUnitRepo,
		
	}
}
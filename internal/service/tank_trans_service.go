package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type TankTransService interface {
	
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


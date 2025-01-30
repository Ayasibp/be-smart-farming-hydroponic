package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type GrowthHistService interface {
	
}

type growthHistService struct {
	growthHistRepo repository.GrowthHistRepository
	
}

type GrowthHistServiceConfig struct {
	GrowthHistRepo repository.GrowthHistRepository
}

func NewGrowthHistService(config GrowthHistServiceConfig) GrowthHistService {
	return &growthHistService{
		growthHistRepo: config.GrowthHistRepo,
	}
}

package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type AggregationService interface {
}

type aggregationService struct {
	aggregationRepo repository.AggregationRepository
	farmRepo        repository.FarmRepository
	systemUnitRepo  repository.SystemUnitRepository
}

type AggregationServiceConfig struct {
	AggregatoionRepo repository.AggregationRepository
	FarmRepo         repository.FarmRepository
	SystemUnitRepo   repository.SystemUnitRepository
}

func NewAggregationService(config AggregationServiceConfig) AggregationService {
	return &aggregationService{
		aggregationRepo: config.AggregatoionRepo,
		farmRepo:        config.FarmRepo,
		systemUnitRepo:  config.SystemUnitRepo,
	}
}

package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
)

type AggregationService interface {
	CreateBatchGrowthHistMonthlyAggregation() (bool, error)
	CreatePrevMonthAggregation() (bool, error)
}

type aggregationService struct {
	aggregationRepo repository.AggregationRepository
	farmRepo        repository.FarmRepository
	systemUnitRepo  repository.SystemUnitRepository
	growthHistRepo  repository.GrowthHistRepository
}

type AggregationServiceConfig struct {
	AggregatoionRepo repository.AggregationRepository
	FarmRepo         repository.FarmRepository
	SystemUnitRepo   repository.SystemUnitRepository
	GrowthHistRepo   repository.GrowthHistRepository
}

func NewAggregationService(config AggregationServiceConfig) AggregationService {
	return &aggregationService{
		aggregationRepo: config.AggregatoionRepo,
		farmRepo:        config.FarmRepo,
		systemUnitRepo:  config.SystemUnitRepo,
		growthHistRepo:  config.GrowthHistRepo,
	}
}

func (s *aggregationService) CreateBatchGrowthHistMonthlyAggregation() (bool, error) {
	logger.Info("aggregationService", "Starting batch growth history monthly aggregation")

	// Fetch aggregated values
	aggregatesVal, err := s.growthHistRepo.GetMonthlyAggregation()
	if err != nil {
		logger.Error("aggregationService", "Failed to fetch monthly aggregation", "error", err)
		return false, err
	}

	var batchValues string
	for _, val := range aggregatesVal {
		for key, value := range val.AggregatedValues {
			batchValues += fmt.Sprintf(
				"('%s','%s','growth-hist',%.2f,'monthly','%s','%d-%d-1','%s'),",
				val.FarmId.String(), val.SystemId.String(), value, key, val.Year, val.Month, time.Now().Format("2006-01-02 15:04:05"),
			)
		}
	}

	batchValues = strings.TrimSuffix(batchValues, ",")
	if batchValues == "" {
		logger.Warn("aggregationService", "No data to insert for batch aggregation")
		return false, nil
	}

	// Insert batch aggregation
	_, err = s.aggregationRepo.CreateBatchAggregation(&batchValues)
	if err != nil {
		logger.Error("aggregationService", "Failed to create batch aggregation", "error", err)
		return false, err
	}

	logger.Info("aggregationService", "Batch growth history monthly aggregation completed successfully")
	return true, nil
}

func (s *aggregationService) CreatePrevMonthAggregation() (bool, error) {
	logger.Info("aggregationService", "Starting previous month aggregation")

	// Fetch aggregated values
	aggregatesVal, err := s.growthHistRepo.GetPrevMonthAggregation()
	if err != nil {
		logger.Error("aggregationService", "Failed to fetch previous month aggregation", "error", err)
		return false, err
	}

	var batchValues string
	for _, val := range aggregatesVal {
		for key, value := range val.AggregatedValues {
			batchValues += fmt.Sprintf(
				"('%s','%s','growth-hist',%.2f,'monthly','%s','%d-%d-1','%s'),",
				val.FarmId.String(), val.SystemId.String(), value, key, val.Year, val.Month, time.Now().Format("2006-01-02 15:04:05"),
			)
		}
	}

	batchValues = strings.TrimSuffix(batchValues, ",")
	if batchValues == "" {
		logger.Warn("aggregationService", "No data to insert for previous month aggregation")
		return false, nil
	}

	// Insert batch aggregation
	_, err = s.aggregationRepo.CreateBatchAggregation(&batchValues)
	if err != nil {
		logger.Error("aggregationService", "Failed to create previous month aggregation", "error", err)
		return false, err
	}

	logger.Info("aggregationService", "Previous month aggregation completed successfully")
	return true, nil
}

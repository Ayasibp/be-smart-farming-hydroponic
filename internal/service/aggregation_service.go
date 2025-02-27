package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
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

	// minus today
	aggregatesVal, err := s.growthHistRepo.GetMonthlyAggregation()
	if err != nil {
		return false, err
	}

	var batchValues string

	for i := 0; i < len(aggregatesVal); i++ {
		val := aggregatesVal[i]
		for key, value := range val.AggregatedValues {
			batchValues = batchValues + "(" + "'" + val.FarmId.String() + "'," + "'" + val.SystemId.String() + "'," + "'growth-hist'" + "," + fmt.Sprintf("%.2f", value) + "," + "'monthly'" + ",'" + key + "'," + "'" + strconv.Itoa(val.Year) + "-" + strconv.Itoa(val.Month) + "-1" + "'" + ",'" + time.Now().Format("2006-01-02 15:04:05") + "')" + ","
		}
	}

	batchValues = strings.TrimSuffix(batchValues, ",")

	_, err = s.aggregationRepo.CreateBatchAggregation(&batchValues)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (s *aggregationService) CreatePrevMonthAggregation() (bool, error) {

	aggregatesVal, err := s.growthHistRepo.GetPrevMonthAggregation()
	if err != nil {
		return false, err
	}

	var batchValues string

	for i := 0; i < len(aggregatesVal); i++ {
		val := aggregatesVal[i]
		for key, value := range val.AggregatedValues {
			batchValues = batchValues + "(" + "'" + val.FarmId.String() + "'," + "'" + val.SystemId.String() + "'," + "'growth-hist'" + "," + fmt.Sprintf("%.2f", value) + "," + "'monthly'" + ",'" + key + "'," + "'" + strconv.Itoa(val.Year) + "-" + strconv.Itoa(val.Month) + "-1" + "'" + ",'" + time.Now().Format("2006-01-02 15:04:05") + "')" + ","
		}
	}

	batchValues = strings.TrimSuffix(batchValues, ",")

	_, err = s.aggregationRepo.CreateBatchAggregation(&batchValues)
	if err != nil {
		return false, nil
	}

	return true, nil
}

package service

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"github.com/google/uuid"
)

type GrowthHistService interface {
	CreateGrowthHist(input *dto.GrowthHist) (*dto.GrowthHistResponse, error)
	GenerateDummyData(input *dto.GrowthHistDummyDataBody) (*dto.GrowthHistResponse, error)
	GetGrowthHistAggregationByFilter(getGrowthFilterBody *dto.GetGrowthFilter) (*dto.GetGrowthAggregationResp, error)
	GetGrowthHistByFilter(getGrowthFilterBody *dto.GetGrowthFilter) (*dto.GetGrowthDataResp, error)
}

type growthHistService struct {
	growthHistRepo  repository.GrowthHistRepository
	farmRepo        repository.FarmRepository
	systemUnitRepo  repository.SystemUnitRepository
	aggregationRepo repository.AggregationRepository
}

type GrowthHistServiceConfig struct {
	GrowthHistRepo  repository.GrowthHistRepository
	FarmRepo        repository.FarmRepository
	SystemUnitRepo  repository.SystemUnitRepository
	AggregationRepo repository.AggregationRepository
}

func NewGrowthHistService(config GrowthHistServiceConfig) GrowthHistService {
	return &growthHistService{
		growthHistRepo:  config.GrowthHistRepo,
		farmRepo:        config.FarmRepo,
		systemUnitRepo:  config.SystemUnitRepo,
		aggregationRepo: config.AggregationRepo,
	}
}

func (s *growthHistService) CreateGrowthHist(input *dto.GrowthHist) (*dto.GrowthHistResponse, error) {
	logger.Info("growthHistService", "Creating Growth History", map[string]string{
		"farmId":   input.FarmId.String(),
		"systemId": input.SystemId.String(),
	})

	farm, err := s.farmRepo.GetFarmById(&model.Farm{
		ID: input.FarmId,
	})
	if err != nil || farm == nil {
		logger.Error("growthHistService", "Invalid Farm ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{
		ID: input.SystemId,
	})
	if err != nil || systemUnit == nil {
		logger.Error("growthHistService", "Invalid System Unit ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidSystemUnitID
	}

	growthHist, err := s.growthHistRepo.CreateGrowthHistory(&model.GrowthHist{
		FarmId:   input.FarmId,
		SystemId: input.SystemId,
		Ppm:      input.Ppm,
		Ph:       input.Ph,
	})
	if err != nil {
		logger.Error("growthHistService", "Error creating new Growth History", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorOnCreatingNewGrowthHist
	}

	respBody := &dto.GrowthHistResponse{
		ID:       growthHist.ID,
		FarmId:   growthHist.FarmId,
		SystemId: growthHist.SystemId,
		Ppm:      growthHist.Ppm,
		Ph:       growthHist.Ph,
	}

	logger.Info("growthHistService", "Growth History created successfully", map[string]string{
		"growthHistId": respBody.ID.String(),
	})
	return respBody, nil
}

func (s *growthHistService) GetGrowthHistAggregationByFilter(getGrowthFilterBody *dto.GetGrowthFilter) (*dto.GetGrowthAggregationResp, error) {
	logger.Info("growthHistService", "Fetching Growth History Aggregation", map[string]string{
		"farmId":   getGrowthFilterBody.FarmId,
		"systemId": getGrowthFilterBody.SystemId,
		"period":   getGrowthFilterBody.Period,
	})

	var aggregateResult *model.GrowthHistAggregate

	farm, err := s.farmRepo.GetFarmById(&model.Farm{
		ID: uuid.MustParse(getGrowthFilterBody.FarmId),
	})
	if err != nil || farm == nil {
		logger.Error("growthHistService", "Invalid Farm ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{
		ID: uuid.MustParse(getGrowthFilterBody.SystemId),
	})
	if err != nil || systemUnit == nil {
		logger.Error("growthHistService", "Invalid System Unit ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidSystemUnitID
	}

	currentDateTime := time.Now()
	var startDate, endDate string

	switch getGrowthFilterBody.Period {
	case "today":
		startDate = currentDateTime.Format("2006-01-02")
		endDate = startDate
	case "last_3_days":
		startDate = currentDateTime.AddDate(0, 0, -3).Format("2006-01-02")
		endDate = currentDateTime.Format("2006-01-02")
	case "last_30_days":
		startDate = currentDateTime.AddDate(0, -1, 0).Format("2006-01-02")
		endDate = currentDateTime.Format("2006-01-02")
	case "custom":
		startDate = getGrowthFilterBody.StartDate.Format("2006-01-02")
		endDate = getGrowthFilterBody.EndDate.Format("2006-01-02")
	}

	aggregateResult, err = s.growthHistRepo.GetAggregateByFilter(getGrowthFilterBody, &startDate, &endDate)
	if err != nil {
		logger.Error("growthHistService", "Error fetching aggregated data", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorOnGettingAggregatedData
	}

	logger.Info("growthHistService", "Successfully fetched Growth History Aggregation", nil)
	return &dto.GetGrowthAggregationResp{
		Period:        getGrowthFilterBody.Period,
		AggregateData: aggregateResult,
	}, nil
}

func (s *growthHistService) GenerateDummyData(input *dto.GrowthHistDummyDataBody) (*dto.GrowthHistResponse, error) {
	start := time.Now()
	logger.Info("growthHistService", "Generating dummy data", map[string]string{
		"farmId":   input.FarmId.String(),
		"systemId": input.SystemId.String(),
	})

	var memStart runtime.MemStats
	runtime.ReadMemStats(&memStart)

	farm, err := s.farmRepo.GetFarmById(&model.Farm{ID: input.FarmId})
	if err != nil || farm == nil {
		logger.Error("growthHistService", "Invalid Farm ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{ID: input.SystemId})
	if err != nil || systemUnit == nil {
		logger.Error("growthHistService", "Invalid System Unit ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidSystemUnitID
	}

	startTime := time.Now().AddDate(-4, 0, 0)
	endTime := time.Now()
	totalJobs := int(endTime.Sub(startTime).Hours())
	numWorkers := int(math.Min(float64(runtime.NumCPU()*2), float64(totalJobs)))
	logger.Info("growthHistService", "Starting dummy data generation", map[string]string{
		"numWorkers": strconv.Itoa(numWorkers),
		"totalJobs":  strconv.Itoa(totalJobs),
	})

	jobs := make(chan time.Time, numWorkers)
	results := make(chan string, numWorkers)

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for t := range jobs {
				farmData := generateRandomFarmData(t)
				record := fmt.Sprintf("('%s','%s',%s,%s,'%s')",
					input.FarmId.String(),
					input.SystemId.String(),
					floatToString(farmData.Ppm),
					floatToString(farmData.Ph),
					farmData.CreatedAt.Format("2006-01-02 15:04:05"),
				)
				results <- record
			}
		}(i)
	}

	go func() {
		for t := startTime; t.Before(endTime); t = t.Add(time.Hour) {
			jobs <- t
		}
		close(jobs)
	}()

	wg.Wait()
	close(results)

	finalBatchValues := strings.Join([]string{}, ",")
	s.growthHistRepo.CreateGrowthHistoryBatch(&finalBatchValues)

	elapsed := time.Since(start)
	logger.Info("growthHistService", "Dummy data generation completed", map[string]string{
		"duration": elapsed.String(),
	})

	return &dto.GrowthHistResponse{
		SystemId: input.SystemId,
		FarmId:   input.FarmId,
	}, nil
}

func (s *growthHistService) GetGrowthHistByFilter(getGrowthFilterBody *dto.GetGrowthFilter) (*dto.GetGrowthDataResp, error) {
	logger.Info("growthHistService", "Fetching Growth History by filter", map[string]string{
		"farmId":   getGrowthFilterBody.FarmId,
		"systemId": getGrowthFilterBody.SystemId,
	})

	farm, err := s.farmRepo.GetFarmById(&model.Farm{
		ID: uuid.MustParse(getGrowthFilterBody.FarmId),
	})
	if err != nil || farm == nil {
		logger.Error("growthHistService", "Invalid Farm ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{
		ID: uuid.MustParse(getGrowthFilterBody.SystemId),
	})
	if err != nil || systemUnit == nil {
		logger.Error("growthHistService", "Invalid System Unit ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.InvalidSystemUnitID
	}

	startDate := getGrowthFilterBody.StartDate.Format("2006-01-02")
	endDate := getGrowthFilterBody.EndDate.Format("2006-01-02")
	aggregateResult, err := s.growthHistRepo.GetDataByFilter(&dto.GetGrowthFilter{
		FarmId:   getGrowthFilterBody.FarmId,
		SystemId: getGrowthFilterBody.SystemId,
	}, &startDate, &endDate)

	if err != nil {
		logger.Error("growthHistService", "Error fetching Growth History data", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("growthHistService", "Successfully fetched Growth History data", nil)
	return &dto.GetGrowthDataResp{
		StartDate: getGrowthFilterBody.StartDate,
		EndDate:   getGrowthFilterBody.EndDate,
		Data:      aggregateResult,
	}, nil
}

func generateRandomFarmData(t time.Time) *model.GrowthHist {
	logger.Info("growthHistService", "Generating random farm data", map[string]string{
		"timestamp": t.String(),
	})
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	return &model.GrowthHist{
		Ppm:       rand.Float64()*1000 + 1, // Random PPM between 1 and 1000
		Ph:        rand.Float64()*14 + 1,   // Random pH between 1 and 14
		CreatedAt: t,
	}
}

func floatToString(input_num float64) string {
	logger.Info("growthHistService", "Converting float to string", map[string]string{
		"value": strconv.FormatFloat(input_num, 'f', 2, 64),
	})
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

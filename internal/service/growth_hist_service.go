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

	farm, err := s.farmRepo.GetFarmById(&model.Farm{
		ID: input.FarmId,
	})
	if err != nil || farm == nil {
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{
		ID: input.SystemId,
	})
	if err != nil || systemUnit == nil {
		return nil, errs.InvalidSystemUnitID
	}

	growthHist, err := s.growthHistRepo.CreateGrowthHistory(&model.GrowthHist{
		FarmId:   input.FarmId,
		SystemId: input.SystemId,
		Ppm:      input.Ppm,
		Ph:       input.Ph,
	})
	if err != nil {
		return nil, errs.ErrorOnCreatingNewGrowthHist
	}

	respBody := &dto.GrowthHistResponse{
		ID:       growthHist.ID,
		FarmId:   growthHist.FarmId,
		SystemId: growthHist.SystemId,
		Ppm:      growthHist.Ppm,
		Ph:       growthHist.Ph,
	}

	return respBody, err
}

func (s *growthHistService) GenerateDummyData(input *dto.GrowthHistDummyDataBody) (*dto.GrowthHistResponse, error) {

	start := time.Now() // Record the start time

	// Track memory usage before execution
	var memStart runtime.MemStats
	runtime.ReadMemStats(&memStart)

	farm, err := s.farmRepo.GetFarmById(&model.Farm{ID: input.FarmId})
	if err != nil || farm == nil {
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{ID: input.SystemId})
	if err != nil || systemUnit == nil {
		return nil, errs.InvalidSystemUnitID
	}

	// Define time range
	startTime := time.Now().AddDate(-4, 0, 0) // 4 years ago
	endTime := time.Now()                     // Current time

	var batchValues strings.Builder
	batchValues.WriteString("") // Initialize the builder

	// Count the total number of jobs before setting numWorkers
	totalJobs := int(endTime.Sub(startTime).Hours()) // One job per hour in 4 years

	// 🔹 Dynamic Worker Allocation (CPU-aware, but limited by total jobs)
	numWorkers := int(math.Min(float64(runtime.NumCPU()*2), float64(totalJobs)))
	fmt.Printf("Using %d workers for %d jobs\n", numWorkers, totalJobs)

	jobs := make(chan time.Time, numWorkers)
	results := make(chan string, numWorkers)

	var wg sync.WaitGroup

	// 🔹 Dynamic Worker Goroutines
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

	// 🔹 Sending jobs dynamically
	go func() {
		for t := startTime; t.Before(endTime); t = t.Add(time.Hour) {
			jobs <- t
		}
		close(jobs) // Close the jobs channel after all jobs are sent
	}()

	// 🔹 Collecting results
	var resultWg sync.WaitGroup
	resultWg.Add(1)
	go func() {
		defer resultWg.Done()
		for result := range results {
			batchValues.WriteString(result + ",")
		}
	}()

	// Wait for all workers to complete
	wg.Wait()
	close(results) // Close results channel only after workers finish

	// Wait for the result collector to complete
	resultWg.Wait()

	finalBatchValues := strings.TrimSuffix(batchValues.String(), ",")

	startDb := time.Now()
	// Store batch data in DB
	s.growthHistRepo.CreateGrowthHistoryBatch(&finalBatchValues)
	fmt.Printf("⏳ DB Insert Time: %v\n", time.Since(startDb))

	// Execution time tracking
	elapsed := time.Since(start)

	// 🔹 Memory Usage Tracking
	var memEnd runtime.MemStats
	runtime.ReadMemStats(&memEnd)
	memUsed := float64(memEnd.Alloc-memStart.Alloc) / (1024 * 1024) // Convert to MB

	fmt.Printf("Execution time: %s\n", elapsed)
	fmt.Printf("Memory used: %.2f MB\n", memUsed)

	return &dto.GrowthHistResponse{
		SystemId: input.SystemId,
		FarmId:   input.FarmId,
	}, nil

}

func (s *growthHistService) GetGrowthHistAggregationByFilter(getGrowthFilterBody *dto.GetGrowthFilter) (*dto.GetGrowthAggregationResp, error) {

	var aggregateResult *model.GrowthHistAggregate

	farm, err := s.farmRepo.GetFarmById(&model.Farm{
		ID: uuid.MustParse(getGrowthFilterBody.FarmId),
	})
	if err != nil || farm == nil {
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{
		ID: uuid.MustParse(getGrowthFilterBody.SystemId),
	})
	if err != nil || systemUnit == nil {
		return nil, errs.InvalidSystemUnitID
	}
	currentDateTime := time.Now()
	if getGrowthFilterBody.Period == "today" {
		currentDate := currentDateTime.Format("2006-01-02")
		aggregateResult, err = s.growthHistRepo.GetAggregateByFilter(&dto.GetGrowthFilter{
			FarmId:   getGrowthFilterBody.FarmId,
			SystemId: getGrowthFilterBody.SystemId,
		}, &currentDate, &currentDate)
		if err != nil {
			return nil, errs.ErrorOnGettingAggregatedData
		}
	}
	if getGrowthFilterBody.Period == "last_3_days" {
		startDate := currentDateTime.AddDate(0, 0, -3).Format("2006-01-02")
		endDate := currentDateTime.Format("2006-01-02")
		aggregateResult, err = s.growthHistRepo.GetAggregateByFilter(&dto.GetGrowthFilter{
			FarmId:   getGrowthFilterBody.FarmId,
			SystemId: getGrowthFilterBody.SystemId,
		}, &startDate, &endDate)
		if err != nil {
			return nil, errs.ErrorOnGettingAggregatedData
		}
	}
	if getGrowthFilterBody.Period == "last_30_days" {
		startDate := currentDateTime.AddDate(0, -1, 0).Format("2006-01-02")
		endDate := currentDateTime.Format("2006-01-02")
		aggregateResult, err = s.growthHistRepo.GetAggregateByFilter(&dto.GetGrowthFilter{
			FarmId:   getGrowthFilterBody.FarmId,
			SystemId: getGrowthFilterBody.SystemId,
		}, &startDate, &endDate)
		if err != nil {
			return nil, errs.ErrorOnGettingAggregatedData
		}
	}
	if getGrowthFilterBody.Period == "custom" {
		farmId, err := uuid.Parse(getGrowthFilterBody.FarmId)
		if err != nil {
			return nil, errs.ErrorOnParsingStringToUUID
		}
		systemId, err := uuid.Parse(getGrowthFilterBody.SystemId)
		if err != nil {
			return nil, errs.ErrorOnParsingStringToUUID
		}
		startDate := getGrowthFilterBody.StartDate.Format("2006-01-02")
		endDate := getGrowthFilterBody.EndDate.Format("2006-01-02")
		aggregatedTableResult, err := s.aggregationRepo.GetAggregatedDataByFilter(&model.Aggregation{
			FarmId: farmId, SystemId: systemId,
		}, &startDate, &endDate)
		if err != nil {
			return nil, errs.ErrorOnGettingAggregatedData
		}
		for _, p := range aggregatedTableResult {
			fmt.Printf("Activity: %s, Value: %f\n", p.Activity, p.Value)
		}
	}

	return &dto.GetGrowthAggregationResp{
		Period:        getGrowthFilterBody.Period,
		AggregateData: aggregateResult,
	}, nil
}

func (s *growthHistService) GetGrowthHistByFilter(getGrowthFilterBody *dto.GetGrowthFilter) (*dto.GetGrowthDataResp, error) {

	farm, err := s.farmRepo.GetFarmById(&model.Farm{
		ID: uuid.MustParse(getGrowthFilterBody.FarmId),
	})
	if err != nil || farm == nil {
		return nil, errs.InvalidFarmID
	}

	systemUnit, err := s.systemUnitRepo.GetSystemUnitById(&model.SystemUnit{
		ID: uuid.MustParse(getGrowthFilterBody.SystemId),
	})
	if err != nil || systemUnit == nil {
		return nil, errs.InvalidSystemUnitID
	}

	startDate := getGrowthFilterBody.StartDate.Format("2006-01-02")
	endDate := getGrowthFilterBody.EndDate.Format("2006-01-02")
	aggregateResult, err := s.growthHistRepo.GetDataByFilter(&dto.GetGrowthFilter{
		FarmId:   getGrowthFilterBody.FarmId,
		SystemId: getGrowthFilterBody.SystemId,
	}, &startDate, &endDate)

	if err != nil {
		return nil, err
	}

	return &dto.GetGrowthDataResp{
		StartDate: getGrowthFilterBody.StartDate,
		EndDate:   getGrowthFilterBody.EndDate,
		Data:      aggregateResult,
	}, nil
}

func generateRandomFarmData(t time.Time) *model.GrowthHist {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	return &model.GrowthHist{
		Ppm:       rand.Float64()*1000 + 1, // Random PPM between 1 and 1000
		Ph:        rand.Float64()*14 + 1,   // Random pH between 1 and 14
		CreatedAt: t,
	}
}
func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

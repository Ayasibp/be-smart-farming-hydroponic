package service

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type GrowthHistService interface {
	CreateGrowthHist(input *dto.GrowthHist) (*dto.GrowthHistResponse, error)
	GenerateDummyData(input *dto.GrowthHistDummyDataBody) (*dto.GrowthHistResponse, error)
	GetGrowthHistByFilter(getGrowthFilterBody *dto.GetGrowthFilter) (*dto.GetGrowthFilter, error)
}

type growthHistService struct {
	growthHistRepo repository.GrowthHistRepository
	farmRepo       repository.FarmRepository
	systemUnitRepo repository.SystemUnitRepository
}

type GrowthHistServiceConfig struct {
	GrowthHistRepo repository.GrowthHistRepository
	FarmRepo       repository.FarmRepository
	SystemUnitRepo repository.SystemUnitRepository
}

func NewGrowthHistService(config GrowthHistServiceConfig) GrowthHistService {
	return &growthHistService{
		growthHistRepo: config.GrowthHistRepo,
		farmRepo:       config.FarmRepo,
		systemUnitRepo: config.SystemUnitRepo,
	}
}

func (s growthHistService) CreateGrowthHist(input *dto.GrowthHist) (*dto.GrowthHistResponse, error) {

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

func (s growthHistService) GenerateDummyData(input *dto.GrowthHistDummyDataBody) (*dto.GrowthHistResponse, error) {

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

	var batchValues string
	// Define the start and end time for the 2-year range
	startTime := time.Now().AddDate(-4, 0, 0) // 2 years ago
	endTime := time.Now()                     // Current time

	// Loop through every hour in the 2-year range
	for t := startTime; t.Before(endTime); t = t.Add(time.Hour) {
		// Generate random farm data for the current hour
		farmData := generateRandomFarmData(t)

		//(farm_id, system_id, ppm, ph, created_at)
		batchValues = batchValues + "(" + "'" + input.FarmId.String() + "'," + "'" + input.SystemId.String() + "'," + FloatToString(farmData.Ppm) + "," + FloatToString(farmData.Ph) + ",'" + farmData.CreatedAt.Format("2006-01-02 15:04:05") + "')" + ","
	}

	batchValues = strings.TrimSuffix(batchValues, ",")

	// fmt.Println(batchValues)

	s.growthHistRepo.CreateGrowthHistoryBatch(&batchValues)

	return &dto.GrowthHistResponse{
		SystemId: input.SystemId,
		FarmId:   input.FarmId,
	}, nil

}

func (s growthHistService) GetGrowthHistByFilter(getGrowthFilterBody *dto.GetGrowthFilter) (*dto.GetGrowthFilter, error) {
	return getGrowthFilterBody, nil
}

func generateRandomFarmData(t time.Time) *model.GrowthHist {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	return &model.GrowthHist{
		Ppm:       rand.Float64()*1000 + 1, // Random PPM between 1 and 1000
		Ph:        rand.Float64()*14 + 1,   // Random pH between 1 and 14
		CreatedAt: t,
	}
}
func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

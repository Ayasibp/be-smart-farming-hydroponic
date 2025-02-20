package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"gorm.io/gorm"
)

type GrowthHistRepository interface {
	CreateGrowthHistory(inputModel *model.GrowthHist) (*model.GrowthHist, error)
	CreateGrowthHistoryBatch(values *string) (int, error)
	GetAggregateByFilter(inputModel *dto.GetGrowthFilter, startDate *string, endDate *string) (*model.GrowthHistAggregate, error)
	GetDataByFilter(inputModel *dto.GetGrowthFilter, startDate *string, endDate *string) ([]*model.GrowthHistFilter, error)
	GetMonthlyAggregation() ([]*model.GrowthHistMonthlyAggregation, error)
	GetPrevMonthAggregation() ([]*model.GrowthHistMonthlyAggregation, error)
}

type growthHistRepository struct {
	db *gorm.DB
}

func NewGrowthHistRepository(db *gorm.DB) GrowthHistRepository {
	return &growthHistRepository{db: db}
}

func (r *growthHistRepository) CreateGrowthHistory(inputModel *model.GrowthHist) (*model.GrowthHist, error) {
	logger.Info("growthHistRepository", "Creating new growth history record")

	sqlScript := `INSERT INTO hydroponic_system.growth_hist(farm_id, system_id, ppm, ph, created_at) 
				  VALUES (?, ?, ?, ?, ?) 
				  RETURNING farm_id, system_id, ppm, ph;`

	res := r.db.Raw(sqlScript, inputModel.FarmId, inputModel.SystemId, inputModel.Ppm, inputModel.Ph, time.Now()).Scan(inputModel)

	if res.Error != nil {
		logger.Error("growthHistRepository", "Failed to create growth history", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("growthHistRepository", "Growth history created successfully", "farm_id", inputModel.FarmId, "system_id", inputModel.SystemId)
	return inputModel, nil
}

func (r *growthHistRepository) CreateGrowthHistoryBatch(values *string) (int, error) {
	logger.Info("growthHistRepository", "Creating batch growth history records")

	var inputModel *model.GrowthHist

	sqlScript := `INSERT INTO hydroponic_system.growth_hist(farm_id, system_id, ppm, ph, created_at) 
				  VALUES ` + *values + ` 
				  RETURNING farm_id, system_id, ppm, ph;`

	res := r.db.Raw(sqlScript).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("growthHistRepository", "Failed to create batch growth history", "error", res.Error)
		return 0, res.Error
	}

	logger.Info("growthHistRepository", "Batch growth history created successfully")
	return 1, nil
}

func (r *growthHistRepository) GetAggregateByFilter(inputModel *dto.GetGrowthFilter, startDate *string, endDate *string) (*model.GrowthHistAggregate, error) {
	logger.Info("growthHistRepository", "Fetching aggregate growth history", "farm_id", inputModel.FarmId, "system_id", inputModel.SystemId)

	outputModel := &model.GrowthHistAggregate{}

	sqlScript := `SELECT
					COALESCE(SUM(ppm),0) as totalPpm,
					COALESCE(SUM(ph),0) as totalPh,
					COALESCE(COUNT(id),0) as totalData,
					COALESCE(MIN(ppm),0) as minPpm,
					COALESCE(MAX(ppm),0) as maxPpm,
					COALESCE(MIN(ph),0) as minPh,
					COALESCE(MAX(ph),0) as maxPh,
					COALESCE(NULLIF(SUM(ppm), 0) / NULLIF(COUNT(id), 0), 0) as avgPpm,
					COALESCE(NULLIF(SUM(ph), 0) / NULLIF(COUNT(id), 0), 0) as avgPh
				  FROM hydroponic_system.growth_hist gh
				  WHERE created_at::date BETWEEN ? AND ?
				  AND farm_id = ?
				  AND system_id = ?;`

	res := r.db.Raw(sqlScript, *startDate, *endDate, inputModel.FarmId, inputModel.SystemId).Scan(outputModel)

	if res.Error != nil {
		logger.Error("growthHistRepository", "Failed to fetch aggregate data", "error", res.Error)
		return nil, res.Error
	}

	return outputModel, nil
}

func (r *growthHistRepository) GetDataByFilter(inputModel *dto.GetGrowthFilter, startDate *string, endDate *string) ([]*model.GrowthHistFilter, error) {
	logger.Info("growthHistRepository", "Fetching filtered growth history data")

	var outputModel []*model.GrowthHistFilter

	sqlScript := `SELECT ppm, ph, created_at
				  FROM hydroponic_system.growth_hist gh
				  WHERE created_at::date BETWEEN ? AND ?
				  AND farm_id = ?
				  AND system_id = ?;`

	res := r.db.Raw(sqlScript, *startDate, *endDate, inputModel.FarmId, inputModel.SystemId).Scan(&outputModel)

	if res.Error != nil {
		logger.Error("growthHistRepository", "Failed to fetch filtered growth history data", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("growthHistRepository", "Filtered growth history data fetched successfully", "count", len(outputModel))
	return outputModel, nil
}

func (r *growthHistRepository) GetMonthlyAggregation() ([]*model.GrowthHistMonthlyAggregation, error) {
	logger.Info("growthHistRepository", "Fetching monthly growth history aggregation")

	var outputModel []*model.GrowthHistMonthlyAggregation

	sqlScript := `SELECT 
					farm_id,
					system_id,
					EXTRACT(YEAR FROM created_at) AS year,
					EXTRACT(MONTH FROM created_at) AS month,
					jsonb_build_object(
						'total_data', COUNT(*),
						'total_ph', ROUND(SUM(ph)::numeric, 2),
						'total_ppm', ROUND(SUM(ppm)::numeric, 2),
						'max_ph', ROUND(MAX(ph)::numeric, 2),
						'min_ph', ROUND(MIN(ph)::numeric, 2),
						'max_ppm', ROUND(MAX(ppm)::numeric, 2),
						'min_ppm', ROUND(MIN(ppm)::numeric, 2)
					) AS aggregated_values
				FROM hydroponic_system.growth_hist gh
				WHERE created_at < DATE_TRUNC('month', CURRENT_DATE)
				GROUP BY 
					farm_id, 
					system_id, 
					EXTRACT(YEAR FROM created_at),
					EXTRACT(MONTH FROM created_at)
				ORDER BY 
					year, 
					month, 
					farm_id, 
					system_id;`

	res := r.db.Raw(sqlScript).Scan(&outputModel)

	if res.Error != nil {
		logger.Error("growthHistRepository", "Failed to fetch monthly aggregation", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("growthHistRepository", "Monthly aggregation fetched successfully", "count", len(outputModel))
	return outputModel, nil
}

func (r *growthHistRepository) GetPrevMonthAggregation() ([]*model.GrowthHistMonthlyAggregation, error) {
	logger.Info("growthHistRepository", "Fetching previous month's growth history aggregation")

	var outputModel []*model.GrowthHistMonthlyAggregation

	sqlScript := `SELECT 
					farm_id,
					system_id,
					EXTRACT(YEAR FROM created_at) AS year,
					EXTRACT(MONTH FROM created_at) AS month,
					jsonb_build_object(
						'avg_ppm', ROUND(AVG(ppm)::numeric, 2),
						'total_data', COUNT(*),
						'total_ph', ROUND(SUM(ph)::numeric, 2),
						'total_ppm', ROUND(SUM(ppm)::numeric, 2),
						'max_ph', ROUND(MAX(ph)::numeric, 2),
						'min_ph', ROUND(MIN(ph)::numeric, 2),
						'max_ppm', ROUND(MAX(ppm)::numeric, 2),
						'min_ppm', ROUND(MIN(ppm)::numeric, 2)
					) AS aggregated_values
				FROM hydroponic_system.growth_hist gh
				WHERE created_at >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
				AND created_at < DATE_TRUNC('month', CURRENT_DATE)
				GROUP BY 
					farm_id, 
					system_id, 
					EXTRACT(YEAR FROM created_at),
					EXTRACT(MONTH FROM created_at)
				ORDER BY 
					year, 
					month, 
					farm_id, 
					system_id;`

	res := r.db.Raw(sqlScript).Scan(&outputModel)

	if res.Error != nil {
		logger.Error("growthHistRepository", "Failed to fetch previous month aggregation", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("growthHistRepository", "Previous month aggregation fetched successfully", "count", len(outputModel))
	return outputModel, nil
}

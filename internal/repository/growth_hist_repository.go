package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type GrowthHistRepository interface {
	CreateGrowthHistory(inputModel *model.GrowthHist) (*model.GrowthHist, error)
	CreateGrowthHistoryBatch(values *string) (int, error)
	GetAggregateByFilter(inputModel *dto.GetGrowthFilter, startDate *string, endDate *string) (*model.GrowthHistAggregate, error)
	GetDataByFilter(inputModel *dto.GetGrowthFilter, startDate *string, endDate *string) ([]*model.GrowthHistFilter, error)
	GetAggregateMonthly() ([]*model.GrowthHistAggregateMonthly, error)
}

type growthHistRepository struct {
	db *gorm.DB
}

func NewGrowthHistRepository(db *gorm.DB) GrowthHistRepository {
	return &growthHistRepository{
		db: db,
	}
}

func (r growthHistRepository) CreateGrowthHistory(inputModel *model.GrowthHist) (*model.GrowthHist, error) {

	sqlScript := `INSERT INTO hydroponic_system.growth_hist(farm_id, system_id, ppm, ph, created_at) 
				VALUES (?,?,?,?,?) 
				RETURNING farm_id, system_id, ppm, ph;`

	res := r.db.Raw(sqlScript, inputModel.FarmId, inputModel.SystemId, inputModel.Ppm, inputModel.Ph, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}

func (r growthHistRepository) CreateGrowthHistoryBatch(values *string) (int, error) {
	var inputModel *model.GrowthHist

	sqlScript := `INSERT INTO hydroponic_system.growth_hist(farm_id, system_id, ppm, ph, created_at) 
				VALUES ` + *values + ` 
				RETURNING farm_id, system_id, ppm, ph;`

	res := r.db.Raw(sqlScript).Scan(&inputModel)

	if res.Error != nil {
		return 0, res.Error
	}
	return 1, nil
}

func (r growthHistRepository) GetAggregateByFilter(inputModel *dto.GetGrowthFilter, startDate *string, endDate *string) (*model.GrowthHistAggregate, error) {
	var outputModel *model.GrowthHistAggregate

	sqlScript := `SELECT
					COALESCE(SUM(ppm),0) as "totalPpm",
					COALESCE(SUM(ph),0) as "totalPh",
					COALESCE(COUNT(id),0) as "totalData",
					COALESCE(MIN(ppm),0) as "minPpm",
					COALESCE(MAX(ppm),0) as "maxPpm",
					COALESCE(MIN(ph),0) as "minPh",
					COALESCE(MAX(ph),0) as "maxPh",
					COALESCE(SUM(ppm)/COUNT(id),0) as "avgPpm",
					COALESCE(SUM(ph)/COUNT(id),0) as "avgPh"
				FROM hydroponic_system.growth_hist gh
				WHERE
					(created_at::date BETWEEN ? AND ?)
					AND farm_id = ?
					AND system_id = ?`

	res := r.db.Raw(sqlScript, *startDate, *endDate, inputModel.FarmId, inputModel.SystemId).Scan(&outputModel)

	if res.Error != nil {
		return outputModel, res.Error
	}

	return outputModel, nil
}
func (r growthHistRepository) GetDataByFilter(inputModel *dto.GetGrowthFilter, startDate *string, endDate *string) ([]*model.GrowthHistFilter, error) {
	var outputModel []*model.GrowthHistFilter

	sqlScript := `SELECT
					ppm,
					ph,
					created_at
				FROM hydroponic_system.growth_hist gh
				WHERE
					(created_at::date BETWEEN ? AND ?)
					AND farm_id = ?
					AND system_id = ?`

	res := r.db.Raw(sqlScript, *startDate, *endDate, inputModel.FarmId, inputModel.SystemId).Scan(&outputModel)

	if res.Error != nil {
		return outputModel, res.Error
	}

	return outputModel, nil
}

func (r growthHistRepository) GetAggregateMonthly() ([]*model.GrowthHistAggregateMonthly, error) {

	var outputModel []*model.GrowthHistAggregateMonthly

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
				WHERE DATE(created_at) < CURRENT_DATE  -- Exclude today's data
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
		return outputModel, res.Error
	}

	return outputModel, nil
}

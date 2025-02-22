package repository

import (
	"strconv"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"gorm.io/gorm"
)

type AggregationRepository interface {
	CreateBatchAggregation(inputValuesString *string) (int, error)
	GetAggregatedDataByFilter(inputModel *model.Aggregation, startDate *string, endDate *string) ([]*model.AggregatedDataByFilter, error)
}

type aggregationRepository struct {
	db *gorm.DB
}

func NewAggregationRepository(db *gorm.DB) AggregationRepository {
	return &aggregationRepository{
		db: db,
	}
}

func (r *aggregationRepository) CreateBatchAggregation(inputValuesString *string) (int, error) {
	logger.Info("aggregationRepository", "Creating batch aggregation", nil)

	var outputModel *int
	sqlScript := `INSERT INTO hydroponic_system.aggregations(farm_id, system_id, name, value, time_range, activity, time, created_at) 
				VALUES ` + *inputValuesString +
		` RETURNING 1;`

	res := r.db.Raw(sqlScript).Scan(&outputModel)

	if res.Error != nil {
		logger.Error("aggregationRepository", "Failed to create batch aggregation", map[string]string{
			"error": res.Error.Error(),
		})
		return 0, res.Error
	}

	logger.Info("aggregationRepository", "Batch aggregation created successfully", nil)
	return 1, nil
}

func (r *aggregationRepository) GetAggregatedDataByFilter(inputModel *model.Aggregation, startDate *string, endDate *string) ([]*model.AggregatedDataByFilter, error) {
	logger.Info("aggregationRepository", "Fetching aggregated data by filter", map[string]string{
		"farmID":    inputModel.FarmId.String(),
		"systemID":  inputModel.SystemId.String(),
		"startDate": *startDate,
		"endDate":   *endDate,
	})

	var outputModel []*model.AggregatedDataByFilter

	sqlScript := `SELECT  
					activity,
					CASE 
						WHEN activity = 'max_ph' THEN MAX(value)
						WHEN activity = 'max_ppm' THEN MAX(value)
						WHEN activity = 'min_ppm' THEN MIN(value)
						WHEN activity = 'min_ph' THEN MIN(value)
						WHEN activity = 'total_ph' THEN SUM(value)
						WHEN activity = 'total_ppm' THEN SUM(value)
						WHEN activity = 'total_data' THEN SUM(value)
					END AS value
				FROM hydroponic_system.aggregations
					WHERE "name" = 'growth-hist'
						AND time_range = 'monthly'
						AND ("time"::date BETWEEN ? AND ?)
						AND farm_id = ?
						AND system_id = ?
				GROUP BY activity`

	res := r.db.Raw(sqlScript, *startDate, *endDate, inputModel.FarmId, inputModel.SystemId).Scan(&outputModel)

	if res.Error != nil {
		logger.Error("aggregationRepository", "Failed to fetch aggregated data", map[string]string{
			"farmID":   inputModel.FarmId.String(),
			"systemID": inputModel.SystemId.String(),
			"error":    res.Error.Error(),
		})
		return outputModel, res.Error
	}

	logger.Info("aggregationRepository", "Aggregated data fetched successfully", map[string]string{
		"count": strconv.Itoa(len(outputModel)),
	})
	return outputModel, nil
}

package repository

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type AggregationRepository interface {
	CreateBatchAggregation(inputValuseString *string) (int, error)
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

func (r *aggregationRepository) CreateBatchAggregation(inputValuseString *string) (int, error) {

	var outputModel *int

	sqlScript := `INSERT INTO hydroponic_system.aggregations(farm_id, system_id, name, value, time_range, activity,time, created_at) 
				VALUES ` + *inputValuseString +
		` RETURNING 1;`

	res := r.db.Raw(sqlScript).Scan(&outputModel)

	if res.Error != nil {
		return 0, res.Error
	}
	return 1, nil
}

func (r *aggregationRepository) GetAggregatedDataByFilter(inputModel *model.Aggregation, startDate *string, endDate *string) ([]*model.AggregatedDataByFilter, error) {
	var outputModel []*model.AggregatedDataByFilter

	sqlScript := `SELECT  
					activity,
					CASE 
						WHEN activity = 'max_ph' THEN MAX(value)
						WHEN activity = 'max_ppm' THEN MAX(value)
						WHEN activity = 'min_ppm'THEN MIN(value)
						WHEN activity  = 'min_ph'THEN MIN(value)
						WHEN activity = 'total_ph'THEN SUM(value)
						WHEN activity = 'total_ppm'THEN SUM(value)
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
		return outputModel, res.Error
	}

	return outputModel, nil
}

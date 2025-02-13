package repository

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type AggregationRepository interface {
	CreateAggregationBatch(inputValuseString *string) (int, error)
	GetAggregationDataByFilter(inputModel *model.Aggregation, startDate *string, endDate *string) ([]*model.Aggregation, error)
}

type aggregationRepository struct {
	db *gorm.DB
}

func NewAggregationRepository(db *gorm.DB) AggregationRepository {
	return &aggregationRepository{
		db: db,
	}
}

func (r *aggregationRepository) CreateAggregationBatch(inputValuseString *string) (int, error) {

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

func (r *aggregationRepository) GetAggregationDataByFilter(inputModel *model.Aggregation, startDate *string, endDate *string) ([]*model.Aggregation, error) {
	var outputModel []*model.Aggregation

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
						AND ("time"::date BETWEEN '2025-01-01'AND '2025-02-01')
						AND farm_id = '803cc3cf-c718-40ed-8411-e5a96b228d2a'
						AND system_id = '3062d904-7560-4a5c-9333-a2325f96167d'
				GROUP BY activity`

	res := r.db.Raw(sqlScript, *startDate, *endDate, inputModel.FarmId, inputModel.SystemId).Scan(&outputModel)

	if res.Error != nil {
		return outputModel, res.Error
	}

	return outputModel, nil
}

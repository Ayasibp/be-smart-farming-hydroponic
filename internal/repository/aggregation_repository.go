package repository

import (
	"gorm.io/gorm"
)

type AggregationRepository interface {
	CreateAggregationBatch(inputValuseString *string) (int, error)
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

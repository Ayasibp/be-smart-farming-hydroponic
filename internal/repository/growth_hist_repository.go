package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type GrowthHistRepository interface {
	CreateGrowthHistory(inputModel *model.GrowthHist) (*model.GrowthHist, error)
	CreateGrowthHistoryBatch(values *string) (int, error)
	GetTodayAggregateByFilter(values *string) (int, error)
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
				VALUES `+*values+` 
				RETURNING farm_id, system_id, ppm, ph;`

	res := r.db.Raw(sqlScript).Scan(&inputModel)

	if res.Error != nil {
		return 0, res.Error
	}
	return 1, nil
}

func (r growthHistRepository) GetTodayAggregateByFilter(values *string) (int, error) {
	var inputModel *model.GrowthHist

	sqlScript := `select 
from growth_hist gh 
where created_at::date = current_date`

	res := r.db.Raw(sqlScript).Scan(&inputModel)

	if res.Error != nil {
		return 0, res.Error
	}
	return 1, nil
}

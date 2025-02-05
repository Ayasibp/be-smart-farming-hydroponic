package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type GrowthHistRepository interface {
	CreateGrowthHistory(inputModel *model.GrowthHist) (*model.GrowthHist, error)
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

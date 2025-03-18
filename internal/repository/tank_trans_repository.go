package repository

import (
	"strconv"
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"gorm.io/gorm"
)

type TankTransRepository interface {
	CreateTankTransaction(inputModel *model.TankTran) (*model.TankTran, error)
}

type tankTransRepository struct {
	db *gorm.DB
}

func NewTankTransRepository(db *gorm.DB) TankTransRepository {
	return &tankTransRepository{db: db}
}

func (r *tankTransRepository) CreateTankTransaction(inputModel *model.TankTran) (*model.TankTran, error) {
	logger.Info("tankTransRepository", "Creating a new tank transaction", map[string]string{
		"farmId":      inputModel.FarmId.String(),
		"systemId":    inputModel.SystemId.String(),
		"waterVolume": strconv.Itoa(inputModel.WaterVolume),
		"aVolume":     strconv.Itoa(inputModel.AVolume),
		"bVolume":     strconv.Itoa(inputModel.BVolume),
	})

	sqlScript := `INSERT INTO hydroponic_system.tank_trans(farm_id, system_id, water_volume, a_volume, b_volume, created_at) 
				  VALUES (?, ?, ?, ?, ?, ?) 
				  RETURNING id, farm_id, system_id, water_volume, a_volume, b_volume, created_at;`

	res := r.db.Raw(sqlScript,
		inputModel.FarmId,
		inputModel.SystemId,
		inputModel.WaterVolume,
		inputModel.AVolume,
		inputModel.BVolume,
		time.Now()).Scan(inputModel)

	if res.Error != nil {
		logger.Error("tankTransRepository", "Failed to create tank transaction", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("tankTransRepository", "Successfully created tank transaction", map[string]string{
		"id": inputModel.ID.String(),
	})
	return inputModel, nil
}

package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"gorm.io/gorm"
)

type SystemLogRepository interface {
	CreateSystemLog(input *model.SystemLog) error
}

type systemLogRepository struct {
	db *gorm.DB
}

func NewSystemLogRepository(db *gorm.DB) SystemLogRepository {
	return &systemLogRepository{
		db: db,
	}
}

func (r *systemLogRepository) CreateSystemLog(input *model.SystemLog) error {
	logger.Info("systemLogRepository", "Creating system log", map[string]string{
		"message": input.Message,
	})

	sqlScript := `INSERT INTO super_admin.system_logs (message, created_at) 
				  VALUES (?,?) 
				  RETURNING id, message;`

	res := r.db.Raw(sqlScript, input.Message, time.Now()).Scan(input)

	if res.Error != nil {
		logger.Error("systemLogRepository", "Failed to create system log", map[string]string{
			"error": res.Error.Error(),
		})
		return res.Error
	}

	logger.Info("systemLogRepository", "System log created successfully", map[string]string{
		"id":      input.ID.String(),
		"message": input.Message,
	})
	return nil
}

package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type SystemLogRepository interface {
	CreateSystemLog(inputModel *model.SystemLog) error
}

type systemLogRepository struct {
	db *gorm.DB
}

func NewSystemLogRepository(db *gorm.DB) SystemLogRepository {
	return &systemLogRepository{
		db: db,
	}
}
func (r systemLogRepository) CreateSystemLog(inputModel *model.SystemLog) error {
	
	res := r.db.Raw("INSERT INTO super_admin.system_logs(message, created_at) VALUES (?,?) RETURNING *;", inputModel.Message, time.Now()).Scan(inputModel)

	if res.Error != nil {
		return res.Error
	}
	return nil

}
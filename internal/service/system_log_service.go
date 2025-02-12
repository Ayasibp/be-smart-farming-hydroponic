package service

import (
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
)

type SystemLogService interface {
	CreateSystemLog(message string) error
}

type systemLogService struct {
	systemLogRepo repository.SystemLogRepository
}

type SystemLogServiceConfig struct {
	SystemLogRepo repository.SystemLogRepository
}

func NewSystemLogService(config SystemLogServiceConfig) SystemLogService {
	return &systemLogService{
		systemLogRepo: config.SystemLogRepo,
	}
}
func (s *systemLogService) CreateSystemLog(message string) error {

	err := s.systemLogRepo.CreateSystemLog(&model.SystemLog{Message: message})
	if err != nil {
		return errs.ErrorCreatingSystemLog
	}

	return nil
}

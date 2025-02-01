package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
)

type UnitIdHandler struct {
	unitIdService    service.UnitIdService
	systemLogService service.SystemLogService
}

type UnitIdHandlerConfig struct {
	UnitIdService    service.UnitIdService
	SystemLogService service.SystemLogService
}

func NewUnitIdHandler(config UnitIdHandlerConfig) *UnitIdHandler {
	return &UnitIdHandler{
		unitIdService:    config.UnitIdService,
		systemLogService: config.SystemLogService,
	}
}

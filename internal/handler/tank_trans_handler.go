package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
)

type TankTransHandler struct {
	tankTransService service.TankTransService
	systemLogService  service.SystemLogService
}

type TankTransHandlerConfig struct {
	TankTransService service.TankTransService
	SystemLogService  service.SystemLogService
}

func NewTankTransHandler(config TankTransHandlerConfig) *TankTransHandler {
	return &TankTransHandler{
		tankTransService: config.TankTransService,
		systemLogService:  config.SystemLogService,
	}
}
package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
)

type GrowthHistHandler struct {
	growthHistService service.GrowthHistService
	systemLogService service.SystemLogService
}

type GrowthHistHandlerConfig struct {
	GrowthHistService service.GrowthHistService
	SystemLogService service.SystemLogService
}

func NewGrowthHistHandler(config GrowthHistHandlerConfig) *GrowthHistHandler {
	return &GrowthHistHandler{
		growthHistService: config.GrowthHistService,
		systemLogService: config.SystemLogService,
	}
}

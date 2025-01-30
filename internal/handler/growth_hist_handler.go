package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
)

type GrowthHistHandler struct {
	growthHistService service.GrowthHistService
}

type GrowthHistHandlerConfig struct {
	GrowthHistService service.GrowthHistService
}

func NewGrowthHistHandler(config GrowthHistHandlerConfig) *GrowthHistHandler {
	return &GrowthHistHandler{
		growthHistService: config.GrowthHistService,
	}
}

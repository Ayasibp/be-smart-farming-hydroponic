package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
)

type FarmHandler struct {
	farmService service.FarmService
}

type FarmHandlerConfig struct {
	FarmService service.FarmService
}

func NewFarmHandler(config FarmHandlerConfig) *FarmHandler {
	return &FarmHandler{
		farmService: config.FarmService,
	}
}

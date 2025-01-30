package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
)

type SystemUnitHandler struct {
	systemUnitService service.SystemUnitService
}

type SystemUnitHandlerConfig struct {
	SystemUnitService service.SystemUnitService
}

func NewSystemUnitHandler(config SystemUnitHandlerConfig) *SystemUnitHandler {
	return &SystemUnitHandler{
		systemUnitService: config.SystemUnitService,
	}
}

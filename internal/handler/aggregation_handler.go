package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
)

type AggregationHandler struct {
	aggregationService service.AggregationService
	systemLogService   service.SystemLogService
}

type AggregationHandlerConfig struct {
	AggregationService service.AggregationService
	SystemLogService   service.SystemLogService
}

func NewAggregationHandler(config AggregationHandlerConfig) *AggregationHandler {
	return &AggregationHandler{
		aggregationService: config.AggregationService,
		systemLogService:   config.SystemLogService,
	}
}

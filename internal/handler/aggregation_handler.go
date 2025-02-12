package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
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

func (h *AggregationHandler) CreateBatchAggregationGrowthHist(c *gin.Context) {

	resp, err := h.aggregationService.CreateBatchGrowthHistMonthlyAggregation()
	if err != nil {
		response.Error(c, 400, err.Error())
	}
	err = h.systemLogService.CreateSystemLog("Create Batch Aggregation: ")
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 201, "Create Aggregation Growth Hist Monthly Success", resp)
}

func (h *AggregationHandler) CreateCurrentMonthAggregationGrowthHist(c *gin.Context) {

	resp, err := h.aggregationService.CreatePrevMonthAggregation()
	if err != nil {
		response.Error(c, 400, err.Error())
	}
	err = h.systemLogService.CreateSystemLog("Create Monthly Aggregation: ")
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 201, "Create Monthly Aggregation Growth Hist Success", resp)
}

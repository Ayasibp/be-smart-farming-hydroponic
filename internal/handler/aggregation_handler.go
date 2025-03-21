package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
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
	logger.Info("aggregationHandler", "Starting batch aggregation growth history process", nil)

	resp, err := h.aggregationService.CreateBatchGrowthHistMonthlyAggregation()
	if err != nil {
		logger.Error("aggregationHandler", "Failed to create batch aggregation growth history", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	err = h.systemLogService.CreateSystemLog("Create Batch Aggregation: ")
	if err != nil {
		logger.Error("aggregationHandler", "Failed to create system log for batch aggregation", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("aggregationHandler", "Successfully created batch aggregation growth history", nil)
	response.JSON(c, 201, "Create Aggregation Growth Hist Monthly Success", resp)
}

func (h *AggregationHandler) CreateCurrentMonthAggregationGrowthHist(c *gin.Context) {
	logger.Info("aggregationHandler", "Starting current month aggregation growth history process", nil)

	resp, err := h.aggregationService.CreatePrevMonthAggregation()
	if err != nil {
		logger.Error("aggregationHandler", "Failed to create current month aggregation growth history", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	err = h.systemLogService.CreateSystemLog("Create Monthly Aggregation: ")
	if err != nil {
		logger.Error("aggregationHandler", "Failed to create system log for monthly aggregation", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("aggregationHandler", "Successfully created current month aggregation growth history", nil)
	response.JSON(c, 201, "Create Monthly Aggregation Growth Hist Success", resp)
}

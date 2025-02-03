package handler

import (
	"encoding/hex"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
)

type GrowthHistHandler struct {
	growthHistService service.GrowthHistService
	systemLogService  service.SystemLogService
}

type GrowthHistHandlerConfig struct {
	GrowthHistService service.GrowthHistService
	SystemLogService  service.SystemLogService
}

func NewGrowthHistHandler(config GrowthHistHandlerConfig) *GrowthHistHandler {
	return &GrowthHistHandler{
		growthHistService: config.GrowthHistService,
		systemLogService:  config.SystemLogService,
	}
}

func (h GrowthHistHandler) CreateGrowthHist(c *gin.Context) {
	var createGrowthHistBody *dto.GrowthHist

	if err := c.ShouldBindJSON(&createGrowthHistBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.growthHistService.CreateGrowthHist(createGrowthHistBody)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	err = h.systemLogService.CreateSystemLog("Create Growth History: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create Growth History Success", resp)
}

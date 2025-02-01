package handler

import (
	"encoding/hex"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
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

func (h UnitIdHandler) CreateUnitId(c *gin.Context) {

	resp, err := h.unitIdService.CreateUnitId()
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	err = h.systemLogService.CreateSystemLog("Create Unit Id: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create Unit Id Success", resp)
}

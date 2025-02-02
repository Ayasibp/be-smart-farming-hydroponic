package handler

import (
	"encoding/hex"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h UnitIdHandler) GetUnitIds(c *gin.Context) {

	resp, err := h.unitIdService.GetUnitIds()
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Unit Id Success", resp)
}

func (h UnitIdHandler) DeleteUnitIdById(c *gin.Context) {

	paramId := c.Param("unitId")
	id, paramErr := uuid.Parse(paramId)
	if paramErr != nil {
		response.Error(c, 400, errs.InvalidUnitKey.Error())
		return
	}
	resp, err := h.unitIdService.DeleteUnitIdbyId(&id)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	err = h.systemLogService.CreateSystemLog("Delete UnitId: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Delete Unit Id Success", resp)
}

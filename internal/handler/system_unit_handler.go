package handler

import (
	"encoding/hex"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
)

type SystemUnitHandler struct {
	systemUnitService service.SystemUnitService
	systemLogService service.SystemLogService
}

type SystemUnitHandlerConfig struct {
	SystemUnitService service.SystemUnitService
	SystemLogService service.SystemLogService
}


func NewSystemUnitHandler(config SystemUnitHandlerConfig) *SystemUnitHandler {
	return &SystemUnitHandler{
		systemUnitService: config.SystemUnitService,
		systemLogService: config.SystemLogService,
	}
}
func (h SystemUnitHandler) CreateSystemUnit(c *gin.Context) {
	var createSystemUnitBody *dto.CreateSystemUnit

	if err := c.ShouldBindJSON(&createSystemUnitBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.systemUnitService.CreateSystemUnit(createSystemUnitBody)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	err = h.systemLogService.CreateSystemLog("Create System Unit: "+ "{ID:"+hex.EncodeToString(resp.ID[:])+"}")
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create System Unit Success", resp)
}
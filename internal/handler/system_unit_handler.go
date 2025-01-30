package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
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

	response.JSON(c, 201, "Create System Unit Success", resp)
}
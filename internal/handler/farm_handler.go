package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h FarmHandler) CreateFarm(c *gin.Context) {
	var createFarmBody *dto.CreateFarm

	if err := c.ShouldBindJSON(&createFarmBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.farmService.CreateFarm(createFarmBody)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create Farm Success", resp)
}

func (h FarmHandler) GetFarms(c *gin.Context) {

	resp, err := h.farmService.GetFarms()
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Farms Success", resp)
}
func (h FarmHandler) GetFarmDetails(c *gin.Context) {

	farmId := c.Param("farmId")
	id, paramErr := uuid.Parse(farmId)
	if paramErr != nil {
		response.Error(c, 400, errs.InvalidFarmID.Error())
		return
	}
	resp, err := h.farmService.GetFarmDetails(&id)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Farm Details Success", resp)
}

func (h FarmHandler) DeleteFarm(c *gin.Context) {

	farmId := c.Param("farmId")
	id, paramErr := uuid.Parse(farmId)
	if paramErr != nil {
		response.Error(c, 400, errs.InvalidFarmIDParam.Error())
		return
	}
	resp, err := h.farmService.DeleteFarm(&id)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Delete Farm Success", resp)
}
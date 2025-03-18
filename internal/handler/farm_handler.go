package handler

import (
	"encoding/hex"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FarmHandler struct {
	farmService      service.FarmService
	systemLogService service.SystemLogService
}

type FarmHandlerConfig struct {
	FarmService      service.FarmService
	SystemLogService service.SystemLogService
}

func NewFarmHandler(config FarmHandlerConfig) *FarmHandler {
	return &FarmHandler{
		farmService:      config.FarmService,
		systemLogService: config.SystemLogService,
	}
}

func (h *FarmHandler) CreateFarm(c *gin.Context) {
	var createFarmBody *dto.CreateFarm

	if err := c.ShouldBindJSON(&createFarmBody); err != nil {
		logger.Error("farmHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	logger.Info("farmHandler", "Starting farm creation process", map[string]string{
		"name": createFarmBody.Name,
	})

	resp, err := h.farmService.CreateFarm(createFarmBody)
	if err != nil {
		logger.Error("farmHandler", "Failed to create farm", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("farmHandler", "Farm created successfully", map[string]string{
		"farmID": hex.EncodeToString(resp.ID[:]),
	})
	err = h.systemLogService.CreateSystemLog("Create Farm: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		logger.Error("farmHandler", "Failed to create system log for farm", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create Farm Success", resp)
}

func (h *FarmHandler) GetFarms(c *gin.Context) {
	logger.Info("farmHandler", "Fetching all farms", nil)

	resp, err := h.farmService.GetFarms()
	if err != nil {
		logger.Error("farmHandler", "Failed to fetch farms", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Farms Success", resp)
}

func (h *FarmHandler) GetFarmDetails(c *gin.Context) {
	farmId := c.Param("farmId")
	id, paramErr := uuid.Parse(farmId)
	if paramErr != nil {
		logger.Error("farmHandler", "Invalid farm ID parameter", map[string]string{
			"farmId": farmId,
		})
		response.Error(c, 400, errs.InvalidFarmID.Error())
		return
	}

	logger.Info("farmHandler", "Fetching farm details", map[string]string{
		"farmID": farmId,
	})

	resp, err := h.farmService.GetFarmDetails(&id)
	if err != nil {
		logger.Error("farmHandler", "Failed to fetch farm details", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Farm Details Success", resp)
}

func (h *FarmHandler) UpdateFarm(c *gin.Context) {
	var updateFarmBody *dto.UpdateFarm

	paramId := c.Param("farmId")
	id, paramErr := uuid.Parse(paramId)
	if paramErr != nil {
		logger.Error("farmHandler", "Invalid farm ID parameter", map[string]string{
			"farmId": paramId,
		})
		response.Error(c, 400, errs.InvalidFarmIDParam.Error())
		return
	}

	if err := c.ShouldBindJSON(&updateFarmBody); err != nil {
		logger.Error("farmHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	logger.Info("farmHandler", "Updating farm", map[string]string{
		"farmID": paramId,
	})

	resp, err := h.farmService.UpdateFarm(&id, updateFarmBody)
	if err != nil {
		logger.Error("farmHandler", "Failed to update farm", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}
	err = h.systemLogService.CreateSystemLog("Update Farm: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		logger.Error("farmHandler", "Failed to create system log for farm update", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Update Farm Success", resp)
}

func (h *FarmHandler) DeleteFarm(c *gin.Context) {
	farmId := c.Param("farmId")
	id, paramErr := uuid.Parse(farmId)
	if paramErr != nil {
		logger.Error("farmHandler", "Invalid farm ID parameter", map[string]string{
			"farmId": farmId,
		})
		response.Error(c, 400, errs.InvalidFarmIDParam.Error())
		return
	}

	logger.Info("farmHandler", "Deleting farm", map[string]string{
		"farmID": farmId,
	})

	resp, err := h.farmService.DeleteFarm(&id)
	if err != nil {
		logger.Error("farmHandler", "Failed to delete farm", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}
	err = h.systemLogService.CreateSystemLog("Delete Farm: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		logger.Error("farmHandler", "Failed to create system log for farm deletion", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Delete Farm Success", resp)
}

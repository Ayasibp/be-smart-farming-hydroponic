package handler

import (
	"bytes"
	"encoding/hex"
	"io"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SystemUnitHandler struct {
	systemUnitService service.SystemUnitService
	systemLogService  service.SystemLogService
}

type SystemUnitHandlerConfig struct {
	SystemUnitService service.SystemUnitService
	SystemLogService  service.SystemLogService
}

func NewSystemUnitHandler(config SystemUnitHandlerConfig) *SystemUnitHandler {
	return &SystemUnitHandler{
		systemUnitService: config.SystemUnitService,
		systemLogService:  config.SystemLogService,
	}
}

func (h *SystemUnitHandler) CreateSystemUnit(c *gin.Context) {
	logger.Info("systemUnitHandler", "Starting CreateSystemUnit process", nil)

	var createSystemUnitBody *dto.CreateSystemUnit
	if err := c.ShouldBindJSON(&createSystemUnitBody); err != nil {
		logger.Error("systemUnitHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.systemUnitService.CreateSystemUnit(createSystemUnitBody)
	if err != nil {
		logger.Error("systemUnitHandler", "Failed to create system unit", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("systemUnitHandler", "Successfully created system unit", map[string]string{
		"SystemUnitID": hex.EncodeToString(resp.ID[:]),
	})

	err = h.systemLogService.CreateSystemLog("Create System Unit: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		logger.Error("systemUnitHandler", "Failed to log system event", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create System Unit Success", resp)
}

func (h *SystemUnitHandler) GetSystemUnits(c *gin.Context) {
	logger.Info("systemUnitHandler", "Starting GetSystemUnits process", nil)

	var systemUnitFilter *dto.SystemUnitFilter
	bodyAsByteArray, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Warn("systemUnitHandler", "Failed to read request body", map[string]string{
			"error": err.Error(),
		})
		systemUnitFilter = nil
	}

	jsonBody := string(bodyAsByteArray)
	if len(jsonBody) == 0 {
		systemUnitFilter = nil
	} else {
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyAsByteArray))
		if err := c.ShouldBindJSON(&systemUnitFilter); err != nil {
			logger.Error("systemUnitHandler", "Invalid request body", map[string]string{
				"error": err.Error(),
			})
			response.Error(c, 400, errs.InvalidRequestBody.Error())
			return
		}
	}

	resp, err := h.systemUnitService.GetSystemUnits(systemUnitFilter)
	if err != nil {
		logger.Error("systemUnitHandler", "Failed to retrieve system units", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("systemUnitHandler", "Successfully retrieved system units", nil)
	response.JSON(c, 200, "Get System Units Success", resp)
}

func (h *SystemUnitHandler) UpdateSystemUnit(c *gin.Context) {
	logger.Info("systemUnitHandler", "Starting UpdateSystemUnit process", nil)

	var updateSystemUnitBody *dto.CreateSystemUnit
	paramId := c.Param("systemId")
	id, paramErr := uuid.Parse(paramId)
	if paramErr != nil {
		logger.Error("systemUnitHandler", "Invalid system unit ID parameter", map[string]string{
			"error": paramErr.Error(),
		})
		response.Error(c, 400, errs.InvalidSystemUnitIDParam.Error())
		return
	}

	if err := c.ShouldBindJSON(&updateSystemUnitBody); err != nil {
		logger.Error("systemUnitHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.systemUnitService.UpdateSystemUnit(&id, updateSystemUnitBody)
	if err != nil {
		logger.Error("systemUnitHandler", "Failed to update system unit", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("systemUnitHandler", "Successfully updated system unit", map[string]string{
		"SystemUnitID": hex.EncodeToString(resp.ID[:]),
	})

	err = h.systemLogService.CreateSystemLog("Update SystemUnit: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		logger.Error("systemUnitHandler", "Failed to log system event", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Update System Unit Success", resp)
}

func (h *SystemUnitHandler) DeleteSystemIdById(c *gin.Context) {
	logger.Info("systemUnitHandler", "Starting DeleteSystemIdById process", nil)

	paramId := c.Param("systemId")
	id, paramErr := uuid.Parse(paramId)
	if paramErr != nil {
		logger.Error("systemUnitHandler", "Invalid system unit ID parameter", map[string]string{
			"error": paramErr.Error(),
		})
		response.Error(c, 400, errs.InvalidSystemUnitID.Error())
		return
	}

	resp, err := h.systemUnitService.DeleteSystemUnitById(&id)
	if err != nil {
		logger.Error("systemUnitHandler", "Failed to delete system unit", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("systemUnitHandler", "Successfully deleted system unit", map[string]string{
		"SystemUnitID": hex.EncodeToString(resp.ID[:]),
	})

	err = h.systemLogService.CreateSystemLog("Delete SystemId: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		logger.Error("systemUnitHandler", "Failed to log system event", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Delete System Id Success", resp)
}

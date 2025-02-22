package handler

import (
	"encoding/hex"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
)

type TankTransHandler struct {
	tankTransService service.TankTransService
	systemLogService service.SystemLogService
}

type TankTransHandlerConfig struct {
	TankTransService service.TankTransService
	SystemLogService service.SystemLogService
}

func NewTankTransHandler(config TankTransHandlerConfig) *TankTransHandler {
	return &TankTransHandler{
		tankTransService: config.TankTransService,
		systemLogService: config.SystemLogService,
	}
}

func (h *TankTransHandler) CreateTankTransaction(c *gin.Context) {
	logger.Info("tankTransHandler", "Starting CreateTankTransaction process", nil)

	var createTankTransBody *dto.TankTransaction
	if err := c.ShouldBindJSON(&createTankTransBody); err != nil {
		logger.Error("tankTransHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.tankTransService.CreateTankTrans(createTankTransBody)
	if err != nil {
		logger.Error("tankTransHandler", "Failed to create tank transaction", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("tankTransHandler", "Successfully created tank transaction", map[string]string{
		"TransactionID": hex.EncodeToString(resp.ID[:]),
	})

	err = h.systemLogService.CreateSystemLog("Create Tank Transaction: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		logger.Error("tankTransHandler", "Failed to log system event", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create Tank Transaction Success", resp)
}

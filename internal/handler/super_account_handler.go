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

type SuperAccountHandler struct {
	superAccountService service.SuperAccountService
	systemLogService    service.SystemLogService
}

type SuperAccountHandlerConfig struct {
	SuperAccountService service.SuperAccountService
	SystemLogService    service.SystemLogService
}

func NewSuperAccountHandler(config SuperAccountHandlerConfig) *SuperAccountHandler {
	return &SuperAccountHandler{
		superAccountService: config.SuperAccountService,
		systemLogService:    config.SystemLogService,
	}
}

func (h *SuperAccountHandler) CreateSuperUser(c *gin.Context) {
	logger.Info("superAccountHandler", "Starting CreateSuperUser process")

	var registerSuperAccountBody *dto.RegisterSuperUserBody
	if err := c.ShouldBindJSON(&registerSuperAccountBody); err != nil {
		logger.Error("superAccountHandler", "Invalid request body", "error", err)
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.superAccountService.CreateSuperUser(registerSuperAccountBody)
	if err != nil {
		logger.Error("superAccountHandler", "Failed to create super user", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("superAccountHandler", "Successfully created super user", "UserID", hex.EncodeToString(resp.UserID[:]))

	err = h.systemLogService.CreateSystemLog("Create Super User: " + "{ID:" + hex.EncodeToString(resp.UserID[:]) + "}")
	if err != nil {
		logger.Error("superAccountHandler", "Failed to log system event", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Register Success", resp)
}

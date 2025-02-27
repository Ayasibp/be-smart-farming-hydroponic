package handler

import (
	"encoding/hex"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountService   service.AccountService
	systemLogService service.SystemLogService
}

type AccountHandlerConfig struct {
	AccountService   service.AccountService
	SystemLogService service.SystemLogService
}

func NewAccountHandler(config AccountHandlerConfig) *AccountHandler {
	return &AccountHandler{
		accountService:   config.AccountService,
		systemLogService: config.SystemLogService,
	}
}

func (h *AccountHandler) CreateUser(c *gin.Context) {
	var registerBody *dto.RegisterBody

	if err := c.ShouldBindJSON(&registerBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.accountService.SignUp(registerBody)

	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	err = h.systemLogService.CreateSystemLog("Create Account: " + "{ID:" + hex.EncodeToString(resp.UserID[:]) + "}")
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	err = h.systemLogService.CreateSystemLog("Create Profile: " + "{ID:" + hex.EncodeToString(resp.ProfileResponse.ID[:]) + "}")
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 201, "Register Success", resp)
}

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

type AccountHandler struct {
	accountService service.AccountService
}

type AccountHandlerConfig struct {
	AccountService service.AccountService
}

func NewAccountHandler(config AccountHandlerConfig) *AccountHandler {
	return &AccountHandler{
		accountService: config.AccountService,
	}
}

func (h *AccountHandler) CreateUser(c *gin.Context) {
	var registerBody *dto.RegisterBody

	if err := c.ShouldBindJSON(&registerBody); err != nil {
		logger.Error("accountHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	logger.Info("accountHandler", "Starting SignUp process", map[string]string{
		"username": registerBody.UserName,
		"email":    registerBody.Email,
	})

	resp, err := h.accountService.SignUp(registerBody)
	if err != nil {
		logger.Error("accountHandler", "SignUp failed", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("accountHandler", "SignUp successful", map[string]string{
		"userID": hex.EncodeToString(resp.UserID[:]),
	})

	response.JSON(c, 201, "Register Success", resp)
}

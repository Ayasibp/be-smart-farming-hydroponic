package handler

import (
	"encoding/hex"
	"net/http"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/tokenprovider"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountService service.AccountService
	tokenProvider  tokenprovider.JWTTokenProvider
}

type AccountHandlerConfig struct {
	AccountService service.AccountService
	TokenProvider  tokenprovider.JWTTokenProvider
}

func NewAccountHandler(config AccountHandlerConfig) *AccountHandler {
	return &AccountHandler{
		accountService: config.AccountService,
		tokenProvider:  config.TokenProvider,
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

func (h *AccountHandler) Login(c *gin.Context) {
	var loginBody dto.LoginBody

	if err := c.ShouldBindJSON(&loginBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.accountService.Login(&loginBody)
	if err != nil {
		logger.Error("AccountHandler", "Failed to login", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("refresh-token", resp.RefreshToken, 3600*24*30, "", "", true, true)
	c.SetCookie("access-token", resp.AccesToken, 3600*24*30, "", "/", true, true)

	response.JSON(c, 200, "Login success", resp)
}

func (h *AccountHandler) Refresh(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	refreshToken, err := h.tokenProvider.ExtractToken(authHeader)
	if err != nil {
		logger.Error("AccountHandler", "Failed to extract refresh token", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	token, err := h.tokenProvider.RenewAccessToken(refreshToken)
	if err != nil {
		logger.Error("AccountHandler Refresh", "Failed to renew access token", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Renew Access Token Success", token)
}

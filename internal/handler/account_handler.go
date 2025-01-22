package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
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

func (h AccountHandler) CreateUser(c *gin.Context) {
	// var registerBody dto.RegisterBody

	// if err := c.ShouldBindJSON(&registerBody); err != nil {
	// 	response.Error(c, 400, errs.InvalidRequestBody.Error())
	// 	return
	// }

	// resp, err := h.authService.CreateUser(registerBody)

	// if err != nil {
	// 	if errors.Is(err, errs.UsernameAlreadyUsed) ||
	// 		errors.Is(err, errs.PasswordContainUsername) {
	// 		response.Error(c, 400, err.Error())
	// 		return
	// 	}

	// 	response.UnknownError(c, err)
	// 	return
	// }

	response.JSON(c, 201, "Register Success", "")
}

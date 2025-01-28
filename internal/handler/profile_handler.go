package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService service.ProfileService
}

type ProfileHandlerConfig struct {
	ProfileService service.ProfileService
}

func NewProfileHandler(config ProfileHandlerConfig) *ProfileHandler {
	return &ProfileHandler{
		profileService: config.ProfileService,
	}
}

func (h ProfileHandler) CreateProfile(c *gin.Context) {
	var createProfileBody dto.CreateProfile

	if err := c.ShouldBindJSON(&createProfileBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.profileService.CreateProfile(createProfileBody)
	if err != nil {
		response.Error(c, 400, err.Error())
	}

	response.JSON(c, 201, "Register Success", resp)
}

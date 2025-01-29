package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var createProfileBody *dto.CreateProfile

	if err := c.ShouldBindJSON(&createProfileBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.profileService.CreateProfile(createProfileBody)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Register Success", resp)
}

func (h ProfileHandler) GetProfileDetails(c *gin.Context) {

	paramId := c.Param("profileId")
	id, paramErr := uuid.Parse(paramId)
	if paramErr != nil {
		response.Error(c, 400, errs.InvalidProfileIDParam.Error())
		return
	}
	resp, err := h.profileService.GetProfileDetails(&id)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Profile Details Success", resp)
}

func (h ProfileHandler) GetProfiles(c *gin.Context) {

	resp, err := h.profileService.GetProfiles()
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Profiles Success", resp)
}

func (h ProfileHandler) DeleteProfile(c *gin.Context) {

	paramId := c.Param("profileId")
	id, paramErr := uuid.Parse(paramId)
	if paramErr != nil {
		response.Error(c, 400, errs.InvalidProfileIDParam.Error())
		return
	}
	resp, err := h.profileService.DeleteProfile(&id)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Delete Success", resp)
}

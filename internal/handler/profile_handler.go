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

type ProfileHandler struct {
	profileService   service.ProfileService
	systemLogService service.SystemLogService
}

type ProfileHandlerConfig struct {
	ProfileService   service.ProfileService
	SystemLogService service.SystemLogService
}

func NewProfileHandler(config ProfileHandlerConfig) *ProfileHandler {
	return &ProfileHandler{
		profileService:   config.ProfileService,
		systemLogService: config.SystemLogService,
	}
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	logger.Info("profileHandler", "Starting CreateProfile process")
	var createProfileBody *dto.CreateProfile

	if err := c.ShouldBindJSON(&createProfileBody); err != nil {
		logger.Error("profileHandler", "Invalid request body", "error", err)
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.profileService.CreateProfile(createProfileBody)
	if err != nil {
		logger.Error("profileHandler", "Error creating profile", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	h.systemLogService.CreateSystemLog("Create Profile: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	logger.Info("profileHandler", "Profile created successfully", "profileID", hex.EncodeToString(resp.ID[:]))

	response.JSON(c, 201, "Create Profile Success", resp)
}

func (h *ProfileHandler) GetProfileDetails(c *gin.Context) {
	logger.Info("profileHandler", "Fetching profile details")
	paramId := c.Param("profileId")
	id, paramErr := uuid.Parse(paramId)
	if paramErr != nil {
		logger.Error("profileHandler", "Invalid profile ID parameter", "error", paramErr)
		response.Error(c, 400, errs.InvalidProfileIDParam.Error())
		return
	}
	resp, err := h.profileService.GetProfileDetails(&id)
	if err != nil {
		logger.Error("profileHandler", "Error fetching profile details", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Profile Details Success", resp)
}

func (h *ProfileHandler) GetProfiles(c *gin.Context) {
	logger.Info("profileHandler", "Fetching all profiles")
	resp, err := h.profileService.GetProfiles()
	if err != nil {
		logger.Error("profileHandler", "Error fetching profiles", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Profiles Success", resp)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	logger.Info("profileHandler", "Updating profile")
	var updateProfileBody *dto.UpdateProfile

	paramId := c.Param("profileId")
	id, paramErr := uuid.Parse(paramId)
	if paramErr != nil {
		logger.Error("profileHandler", "Invalid profile ID parameter", "error", paramErr)
		response.Error(c, 400, errs.InvalidProfileIDParam.Error())
		return
	}

	if err := c.ShouldBindJSON(&updateProfileBody); err != nil {
		logger.Error("profileHandler", "Invalid request body", "error", err)
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.profileService.UpdateProfile(&id, updateProfileBody)
	if err != nil {
		logger.Error("profileHandler", "Error updating profile", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	h.systemLogService.CreateSystemLog("Update Profile: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	logger.Info("profileHandler", "Profile updated successfully", "profileID", hex.EncodeToString(resp.ID[:]))

	response.JSON(c, 201, "Update Profile Success", resp)
}

func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	logger.Info("profileHandler", "Deleting profile")
	paramId := c.Param("profileId")
	id, paramErr := uuid.Parse(paramId)
	if paramErr != nil {
		logger.Error("profileHandler", "Invalid profile ID parameter", "error", paramErr)
		response.Error(c, 400, errs.InvalidProfileIDParam.Error())
		return
	}
	resp, err := h.profileService.DeleteProfile(&id)
	if err != nil {
		logger.Error("profileHandler", "Error deleting profile", "error", err)
		response.Error(c, 400, err.Error())
		return
	}

	h.systemLogService.CreateSystemLog("Delete Profile: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	logger.Info("profileHandler", "Profile deleted successfully", "profileID", hex.EncodeToString(resp.ID[:]))

	response.JSON(c, 201, "Delete Profile Success", resp)
}

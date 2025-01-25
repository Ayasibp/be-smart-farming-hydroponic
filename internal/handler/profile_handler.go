package handler

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
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

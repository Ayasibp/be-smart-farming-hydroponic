package routes

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Account *handler.AccountHandler
	Profile *handler.ProfileHandler
}

type Middlewares struct {
	Auth gin.HandlerFunc
}

func Build(srv *gin.Engine, h Handlers, middleware Middlewares) {
	auth := srv.Group("/auth")
	auth.POST("/register", h.Account.CreateUser)

	profile := srv.Group("/profile")
	profile.POST("/create", h.Profile.CreateProfile)
	profile.GET("/:profileId", h.Profile.GetProfileDetails)
	profile.GET("/", h.Profile.GetProfiles)
	profile.DELETE("/:profileId", h.Profile.DeleteProfile)
}

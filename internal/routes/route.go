package routes

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Account *handler.AccountHandler
	Profile *handler.ProfileHandler
	Farm    *handler.FarmHandler
	SystemUnit *handler.SystemUnitHandler
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
	profile.PUT("/:profileId", h.Profile.UpdateProfile)
	profile.DELETE("/:profileId", h.Profile.DeleteProfile)

	farm := srv.Group("/farm")
	farm.POST("/create", h.Farm.CreateFarm)
	farm.GET("/", h.Farm.GetFarms)
	farm.GET("/:farmId", h.Farm.GetFarmDetails)
	farm.PUT("/:farmId", h.Farm.UpdateFarm)
	farm.DELETE("/:farmId", h.Farm.DeleteFarm)

	systemUnit := srv.Group("/system")
	systemUnit.POST("/create", h.SystemUnit.CreateSystemUnit)
}

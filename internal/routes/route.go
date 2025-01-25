package routes

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Account *handler.AccountHandler
}

type Middlewares struct {
	Auth gin.HandlerFunc
}

func Build(srv *gin.Engine, h Handlers, middleware Middlewares) {
	auth := srv.Group("/auth")
	auth.POST("/register", h.Account.CreateUser)
}

package routes

import "github.com/gin-gonic/gin"

type Handlers struct {
}

type Middlewares struct {
	Auth gin.HandlerFunc
}

func Build(srv *gin.Engine, h Handlers, middleware Middlewares) {

}

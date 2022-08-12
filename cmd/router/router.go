package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{
		handler: handler,
	}
}

func (router *Router) Run() {
	r := gin.Default()
	r.GET("/profile/:id", router.handler.getProfile)
	r.POST("/profile/:id", router.handler.updateProfile)
	r.POST("/screech", router.handler.createScreech)
	r.GET("/screech/:id", router.handler.getScreech)
	r.POST("/screech/:id", router.handler.updateScreech)
	r.GET("/screechlist", router.handler.getScreechList)

	r.Run(":80")
}

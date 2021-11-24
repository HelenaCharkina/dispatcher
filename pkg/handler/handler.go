package handler

import (
	"dispatcher/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		agents := api.Group("/agents")
		{
			agents.GET("/", h.getAllAgents)
			agents.GET("/:id", h.getAgentById)
			agents.POST("/", h.addAgent)
			agents.PUT("/:id", h.updateAgent)
		}
	}
	return router
}

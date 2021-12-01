package handler

import (
	"dispatcher/pkg/service"
	"dispatcher/pkg/settings"
	"fmt"
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
	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.GET("/refresh/:user_id", h.refreshToken)
		auth.GET("/logout/:user_id", h.logout)
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("http://%s:%s", settings.Config.ClientHost, settings.Config.ClientPort))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


package handler

import (
	"dispatcher/pkg/service"
	"dispatcher/pkg/settings"
	"dispatcher/types"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	wsChan  <-chan types.WsChanMessage
	cmdChan chan types.CmdChanMessage
}

func NewHandler(service *service.Service, wsChan <-chan types.WsChanMessage, cmdChan chan types.CmdChanMessage) *Handler {
	return &Handler{
		service: service,
		wsChan:  wsChan,
		cmdChan: cmdChan,
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
			agents.PUT("/", h.updateAgent)
			agents.DELETE("/:id", h.deleteAgent)

			agents.GET("/start/:id", h.start)
			agents.GET("/stop/:id", h.stop)
		}

		statistics := api.Group("/statistics")
		{
			statistics.POST("/", h.getStatisticsByAgent)
		}
	}

	router.GET("/ws", func(c *gin.Context) {
		h.wshandler(c.Writer, c.Request)
	})

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("http://%s:%s", settings.Config.ClientHost, settings.Config.ClientPort))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

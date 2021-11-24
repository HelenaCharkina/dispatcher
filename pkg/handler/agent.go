package handler

import (
	"dispatcher/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) addAgent(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		types.NewErrorResponse(c, http.StatusUnauthorized, "user not found")
		return
	}
}

func (h Handler) getAgentById(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		types.NewErrorResponse(c, http.StatusUnauthorized, "user not found")
		return
	}
}

func (h Handler) updateAgent(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		types.NewErrorResponse(c, http.StatusUnauthorized, "user not found")
		return
	}
}

func (h Handler) getAllAgents(c *gin.Context) {
	fmt.Println("get all agents!!!")
	_, ok := c.Get(userCtx)
	if !ok {
		types.NewErrorResponse(c, http.StatusUnauthorized, "user not found")
		return
	}
}

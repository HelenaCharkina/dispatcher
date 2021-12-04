package handler

import (
	"dispatcher/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) addAgent(c *gin.Context) {
	var agent types.Agent
	if err := c.BindJSON(&agent); err != nil {
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Agent.Add(&agent)
	if err != nil {
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h Handler) getAgentById(c *gin.Context) {

}

func (h Handler) updateAgent(c *gin.Context) {

}

func (h Handler) getAllAgents(c *gin.Context) {
	agents, err := h.service.Agent.GetAll()
	if err != nil {
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, agents)
}

func (h Handler) deleteAgent(c *gin.Context) {

}
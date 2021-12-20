package handler

import (
	"dispatcher/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addAgent(c *gin.Context) {
	var agent types.Agent
	if err := c.BindJSON(&agent); err != nil {
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Agent.Add(&agent)
	if err != nil {
		types.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) getAgentById(c *gin.Context) {

}

func (h *Handler) updateAgent(c *gin.Context) {
	var agent types.Agent
	if err := c.BindJSON(&agent); err != nil {
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Agent.Update(&agent)
	if err != nil {
		types.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) getAllAgents(c *gin.Context) {
	agents, err := h.service.Agent.GetAll()
	if err != nil {
		types.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, agents)
}

func (h *Handler) deleteAgent(c *gin.Context) {
	agentId := c.Param("id")
	if agentId == "" {
		types.NewErrorResponse(c, http.StatusInternalServerError, "agent not found")
		return
	}
	err := h.service.Agent.Delete(agentId)
	if err != nil {
		types.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}


func (h *Handler) start(c *gin.Context) {
	agentId := c.Param("id")
	if agentId == "" {
		types.NewErrorResponse(c, http.StatusInternalServerError, "agent not found")
		return
	}
	h.cmdChan <- types.CmdChanMessage{
		Message: types.START,
		AgentId: agentId,
	}

	c.JSON(http.StatusOK, nil)
}


func (h *Handler) stop(c *gin.Context) {
	agentId := c.Param("id")
	if agentId == "" {
		types.NewErrorResponse(c, http.StatusInternalServerError, "agent not found")
		return
	}
	h.cmdChan <- types.CmdChanMessage{
		Message: types.STOP,
		AgentId: agentId,
	}

	c.JSON(http.StatusOK, nil)
}
package handler

import (
	"dispatcher/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getStatisticsByAgent(c *gin.Context) {

	var params types.StatisticsRequest
	if err := c.BindJSON(&params); err != nil {
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	stats, err := h.service.Statistics.GetByAgentId(&params)
	if err != nil {
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, stats)
}

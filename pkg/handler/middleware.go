package handler

import (
	"dispatcher/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		types.NewErrorResponse(c, http.StatusBadRequest, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != `Bearer` {
		types.NewErrorResponse(c, http.StatusBadRequest, "invalid auth header")
		return
	}

	userId, err := h.service.Authorization.CheckToken(headerParts[1])
	if err != nil {
		types.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

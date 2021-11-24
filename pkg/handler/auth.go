package handler

import (
	"dispatcher/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) signIn(c *gin.Context) {
	var user types.User
	if err := c.BindJSON(&user); err != nil {
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(user.Login, user.Password)
	if err != nil {
		types.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h Handler) signUp(c *gin.Context) {

}

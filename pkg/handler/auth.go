package handler

import (
	"dispatcher/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) login(c *gin.Context) {
	var user types.User
	if err := c.BindJSON(&user); err != nil {
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.service.Authorization.SignIn(user.Login, user.Password)
	if err != nil {
		types.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie("token", response.RefreshToken, 100000, "/auth", "", false, true)
	c.JSON(http.StatusOK, response)
}

func (h Handler) refreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			types.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		types.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId := c.Param("user_id")
	if userId == "" {
		types.NewErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	response, err := h.service.Authorization.RefreshToken(refreshToken, userId)
	if err != nil {
		types.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie("token", response.RefreshToken, 100000, "/auth", "", false, true)
	c.JSON(http.StatusOK, response)
}

func (h Handler) logout(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		types.NewErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

}

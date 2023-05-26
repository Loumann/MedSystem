package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())

		return
	}

	if err := h.r.UserIsExist(input.Username, input.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

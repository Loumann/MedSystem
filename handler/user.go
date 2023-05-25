package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserInput struct {
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic" binding:"required"`
	Snils      string `json:"snils" binding:"required"`
}

func (h *Handler) CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.r.CreateUser(
		input.Name,
		input.Surname,
		input.Patronymic,
		input.Snils,
	); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (h *Handler) GetUsers(c *gin.Context) {
	search := c.Param("search")

	users, err := h.r.GetUsers(search)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatusJSON(200, users)
}

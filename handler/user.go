package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateUserInput struct {
	id         int    `json:"id" binding:"required"`
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

func (h *Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.r.DeleteUser(id); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
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

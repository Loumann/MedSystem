package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"somename/models"
	"strconv"
)

func (h *Handler) WaitAnalyse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())

		return
	}

	user, err := h.r.GetUserByID(id)
	if err != nil {
		fmt.Println("2")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.r.AppendWaitUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		fmt.Println("3")

		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetWaitingUsers(c *gin.Context) {
	users, err := h.r.GetWaitingUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) FulfillWaitingUser(c *gin.Context) {
	var input models.FulfillWaitingUserInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		log.Println("Error fulfillwaitingUser")

		return
	}

	if err := h.r.LinkUserWithAnalyse(input.User, &input.Analyse); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.r.RemoveWaitingUser(input.User); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

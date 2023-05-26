package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetAnalyse(c *gin.Context) {
	// проверка на авторизацию
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		log.Println(id)

		log.Println(err)
		return
	}

	analisys, err := h.r.GetAnalisys(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, analisys)
}

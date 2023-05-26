package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "LoginTemplate.html", nil)
}

func (h *Handler) IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "MainTemplate.html", nil)
}

func (h *Handler) IndexPage11(c *gin.Context) {
	c.HTML(http.StatusOK, "MainTemplate.html", nil)
}

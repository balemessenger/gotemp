package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func NewHttpHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 1})
	return
}

package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"{{ProjectName}}/internal/entity"
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

func (h *Handler) AdminExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello Admin"})
	return
}

func (h *Handler) Example(c *gin.Context) {
	var example entity.Example
	err := c.ShouldBindJSON(&example)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"title": example.Title, "body": example.Body})
	return
}
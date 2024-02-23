package system

import "github.com/gin-gonic/gin"

type SystemHandler struct {}

func(h *SystemHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}
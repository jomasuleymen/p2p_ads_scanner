package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrorMessage(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsernameHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "username",
		"data": username,
	})
}

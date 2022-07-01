package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CekRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("id")
		c.Set("id", id)

		if c.Request.Method == "GET" {
			c.Next()
			return
		}

		response := struct {
			Error_message string `json:"error_message"`
			Error_key     string `json:"error_key"`
		}{
			Error_message: "failed",
			Error_key:     "You are not admin",
		}

		if !CekRoleByID(id) {
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		c.Next()
	}
}

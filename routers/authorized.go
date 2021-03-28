package routers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jkbmdk/kanban-api/pkg/jwt"
)

type unauthorized struct {
	Message string `json:"message"`
}

func authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")
		if _, err := jwt.Verify(token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorized{Message: err.Error()})
			return
		}
		c.Next()
	}
}

package routers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jkbmdk/kanban-api/internal/models"
	"github.com/jkbmdk/kanban-api/pkg/jwt"
	"github.com/jkbmdk/kanban-api/pkg/users"
)

type unauthorized struct {
	Message string `json:"message"`
}

func authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.User
		header := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")
		id, err := jwt.Parse(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorized{Message: err.Error()})
			return
		}
		user, err = users.GetUserById(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorized{Message: "Requested user not found"})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

package auth

import (
	"github.com/jkbmdk/kanban-api/internal/models"
	"github.com/jkbmdk/kanban-api/pkg/users"
)

type RegisterForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Register(form *RegisterForm) (*models.User, error) {
	var u models.User
	u.Email = form.Email
	u.SetPassword(form.Password)
	err := users.StoreUser(&u)
	return &u, err
}

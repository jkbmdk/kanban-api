package auth

import (
    "errors"

    "github.com/jkbmdk/kanban-api/pkg/jwt"
    "github.com/jkbmdk/kanban-api/pkg/users"
)

type GrantAccessForm struct {
    Email    string `form:"email" json:"email" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

type Access struct {
    Token string `json:"access_token"`
}

func GrantAccess(form *GrantAccessForm) (*Access, error) {
    var access Access
    user, err := users.GetUserByEmail(form.Email)
    if err != nil || !user.VerifyPassword(form.Password)  {
        return &access, errors.New("wrong credentials")
    }
    access.Token = jwt.Generate(&user)
    return &access, nil
}
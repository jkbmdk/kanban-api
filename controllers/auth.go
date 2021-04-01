package controllers

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jkbmdk/kanban-api/internal/models"
    "github.com/jkbmdk/kanban-api/pkg/auth"
    "github.com/jkbmdk/kanban-api/pkg/mailer"
)

func Register(c *gin.Context) {
    var form auth.RegisterForm
    var user *models.User

    err := c.BindJSON(&form)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
        return
    }
    user, err = auth.Register(&form)
    if err != nil {
        fmt.Println(err.Error())
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    mail := mailer.Mail{
        Subject:  "Please verify your email",
        Template: "verify",
        From:     "kanban@example.com",
    }
    err = mail.Send(user.Email)
    if err != nil {
        // shall remove user
        fmt.Println(err.Error())
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    c.JSON(http.StatusCreated, user)
}

func GrantAccess(c *gin.Context) {
    var form auth.GrantAccessForm
    err := c.BindJSON(&form)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
        return
    }
    access, err2 := auth.GrantAccess(&form)
    if err2 != nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, err2.Error())
    } else {
        c.JSON(http.StatusCreated, access)
    }
}

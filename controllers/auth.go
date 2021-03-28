package controllers

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jkbmdk/kanban-api/pkg/auth"
)

func Register(c *gin.Context) {
    var form auth.RegisterForm
    err := c.BindJSON(&form)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
        return
    }
    user, err2 := auth.Register(&form)
    if err2 != nil {
        fmt.Println(err2.Error())
        c.AbortWithStatus(http.StatusInternalServerError)
    } else {
        c.JSON(http.StatusCreated, user)
    }
}
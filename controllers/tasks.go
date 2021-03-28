package controllers

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/jkbmdk/kanban-api/internal/models"
    "github.com/jkbmdk/kanban-api/pkg/tasks"
)

func AllTasks(c *gin.Context) {
    t := tasks.GetAllTasks()
    c.JSON(http.StatusOK, t)
}

func ShowTask(c *gin.Context) {
    var task models.Task
    var id int
    var err error

    id, err = parseId(c.Params.ByName("id"))
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    task, err = tasks.GetTaskByID(id)
    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, task)
    }
}

func StoreTask(c *gin.Context) {
    var task models.Task

    err2 := c.BindJSON(&task)
    fmt.Println(err2.Error())

    err := tasks.StoreTask(&task)
    if err != nil {
        fmt.Println(err.Error())
        c.AbortWithStatus(http.StatusInternalServerError)
    } else {
        c.JSON(http.StatusCreated, task)
    }
}

func UpdateTask(c *gin.Context) {
    var task models.Task
    var id int
    var err error

    id, err = parseId(c.Params.ByName("id"))
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    task, err = tasks.GetTaskByID(id)
    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        return
    }

    _ = c.BindJSON(&task)
    err = tasks.UpdateTask(&task)
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
    } else {
        c.JSON(http.StatusOK, task)
    }
}

func DeleteTask(c *gin.Context) {
    var id int
    var err error

    id, err = parseId(c.Params.ByName("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    err = tasks.DeleteTask(id)
    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.Status(http.StatusNoContent)
    }
}

func parseId(id string) (int, error) {
    return strconv.Atoi(id)
}

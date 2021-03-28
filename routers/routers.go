package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/jkbmdk/kanban-api/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    auth := r.Group("/auth")
    {
        auth.POST("register", controllers.Register)
        auth.POST("access", controllers.GrantAccess)
    }
    tasks := r.Group("/tasks", authorized())
    {
        tasks.GET("", controllers.AllTasks)
        tasks.GET(":id", controllers.ShowTask)
        tasks.POST("", controllers.StoreTask)
        tasks.PUT(":id", controllers.UpdateTask)
        tasks.DELETE(":id", controllers.DeleteTask)
    }
    return r
}

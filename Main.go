package main

import (
    "github.com/jkbmdk/kanban-api/internal/database"
    "github.com/jkbmdk/kanban-api/routers"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic("Create and configure .env file using `cp .env .env.example`")
    }
    database.ConnectDataBase()
    r := routers.SetupRouter()
    err = r.Run()
    if err != nil {
        panic(err)
    }
}

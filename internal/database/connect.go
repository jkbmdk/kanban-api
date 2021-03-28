package database

import (
    "github.com/jkbmdk/kanban-api/internal/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
    database, err := gorm.Open(sqlite.Open("kanban.db"))

    if err != nil {
        panic("failed to connect to database")
    }

    err = database.AutoMigrate(&models.User{})
    err = database.AutoMigrate(&models.Task{})

    if err != nil {
        panic("failed to proceed migration")
    }

    DB = database
}

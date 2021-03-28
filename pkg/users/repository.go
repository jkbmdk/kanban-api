package users

import (
    "github.com/jkbmdk/kanban-api/internal/database"
    "github.com/jkbmdk/kanban-api/internal/models"
)

func GetAllUsers() []models.User {
    var users []models.User
    database.DB.Find(&users)
    return users
}

func GetUserByEmail(email string) (models.User, error) {
    var user models.User
    err := database.DB.Where("email = ?", email).First(&user).Error
    return user, err
}

func StoreUser(user *models.User) error {
    return database.DB.Create(user).Error
}

package models

import (
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    Model
    Email    string `json:"email" gorm:"uniqueIndex;not null"`
    Password string `json:"-" gorm:"not null"`
}

func (u *User) SetPassword(password string) {
    p, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
    u.Password = string(p)
}

func (u *User) VerifyPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}

func (u *User) GetID() uint {
    return u.Model.ID
}

func (u *User) GetEmail() string {
    return u.Email
}

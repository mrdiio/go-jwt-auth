package models

import (
	"github.com/mrdiio/go-jwt-auth/response"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `json:"name" gorm:"not null"`
	Username string  `json:"username" gorm:"unique;not null"`
	Email    *string `json:"email" gorm:"unique"`
	Password string  `json:"password" gorm:"not null"`
}

type UserRepo interface {
	Create(user *User) error
	FindAll() []User
	Delete(user *User)
}

type UserService interface {
	Create(user *User) error
	FindAll() []response.UserResponse
	Delete(user *User)
}

func (u *User) ToUserResponse() response.UserResponse {
	return response.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
	}
}

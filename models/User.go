package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrdiio/go-jwt-auth/response"
)

type User struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name             string    `json:"name" gorm:"not null"`
	Username         string    `json:"username" gorm:"unique;not null"`
	Email            string    `json:"email" gorm:"unique;not_null"`
	Password         string    `json:"password" gorm:"not null"`
	VerificationCode string    `json:"verification_code"`
	Verified         bool      `json:"verified" gorm:"not_null;default:false"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type UserRepo interface {
	Create(user *User) error
	FindAll() []User
	Delete(user *User)
	FindByEmail(email string) (*User, error)
}

type UserService interface {
	Create(user *User) error
	FindAll() []response.UserResponse
	Delete(user *User)
	VerifyCredentials(email string, password string) (*response.LoginResponse, error)
}

func (u *User) ToUserResponse() response.UserResponse {
	return response.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
	}
}

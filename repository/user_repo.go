package repository

import (
	"github.com/mrdiio/go-jwt-auth/helper"
	"github.com/mrdiio/go-jwt-auth/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) models.UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
	result := r.db.Create(&user)
	return result.Error
}

func (r *userRepo) FindAll() []models.User {
	var users []models.User
	result := r.db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

func (r *userRepo) Delete(user *models.User) {
	result := r.db.Delete(&user)
	helper.ErrorPanic(result.Error)
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

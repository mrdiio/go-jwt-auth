package service

import (
	"github.com/mrdiio/go-jwt-auth/models"
	"github.com/mrdiio/go-jwt-auth/response"
)

type userService struct {
	UserRepo models.UserRepo
}

func NewUserService(userRepo models.UserRepo) models.UserService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (s *userService) Create(user *models.User) error {
	err := s.UserRepo.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) FindAll() []response.UserResponse {
	result := s.UserRepo.FindAll()
	var users []response.UserResponse

	for _, user := range result {
		users = append(users, response.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return users
}

func (s *userService) Delete(user *models.User) {
	s.UserRepo.Delete(user)
}

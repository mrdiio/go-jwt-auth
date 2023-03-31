package service

import (
	"errors"

	"github.com/mrdiio/go-jwt-auth/config"
	"github.com/mrdiio/go-jwt-auth/helper"
	"github.com/mrdiio/go-jwt-auth/models"
	"github.com/mrdiio/go-jwt-auth/response"
	"github.com/spf13/viper"
)

type userService struct {
	UserRepo models.UserRepo
}

func NewUserService(userRepo models.UserRepo) models.UserService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (s *userService) VerifyCredentials(email string, password string) (*response.LoginResponse, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if !helper.VerifyPassword(password, user.Password) {
		return nil, errors.New("wrong password")
	}

	env := config.LoadEnv()

	// generate jwt token
	accessToken, err := helper.GenerateToken(user.ID, 10, env.AccessTokenSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := helper.GenerateToken(user.ID, 10, viper.GetString("REFRESH_TOKEN_SECRET"))
	if err != nil {
		return nil, err
	}

	res := response.LoginResponse{
		User:         user.ToUserResponse(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &res, nil
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

package services

import (
	"context"
	"fmt"

	"github.com/harshit3011/Insieme/models"
	"github.com/harshit3011/Insieme/repositories"
	"github.com/harshit3011/Insieme/utils"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, user models.RegisterDetails) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	newUser := models.User{
		Name:         user.Name,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: hashedPassword,
		ProfilePic:   user.ProfilePic,
		Bio:          user.Bio,
	}
	if err := s.UserRepo.CreateUser(ctx, newUser); err != nil {
		return fmt.Errorf("Error while creating user: %w", err)
	}
	return nil
}

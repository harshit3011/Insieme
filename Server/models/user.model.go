package models

import "time"

type User struct {
	ID           string
	Name         string
	Username     string
	Email        string
	PasswordHash string
	ProfilePic   *string
	Bio          *string
	CreatedAt    time.Time
	UpdatedAt	time.Time
}

type RegisterDetails struct {
	Name       string  `json:"name" validate:"required"`
	Username   string  `json:"username" validate:"required"`
	Email      string  `json:"email" validate:"required,email"`
	Password   string  `json:"password" validate:"required"`
	ProfilePic *string `json:"profilepic"`
	Bio        *string `json:"bio"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

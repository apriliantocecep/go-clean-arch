package model

import (
	"go-clean-arch/internal/entity"
)

type UserResponse struct {
	User  *entity.User `json:"user,omitempty"`
	Token string       `json:"token,omitempty"`
}

type RegisterUserRequest struct {
	Name     string `json:"name" binding:"required,gte=3,lte=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=8,lte=255"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ValidateUserRequest struct {
	Token string `json:"token" binding:"required"`
}

type GetUserRequest struct {
	ID uint `json:"id" binding:"required"`
}

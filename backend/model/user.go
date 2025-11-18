package model

import "github.com/google/uuid"

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserParam struct {
	ID       uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
}

type LoginAcc struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type ChangePassword struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

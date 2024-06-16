package models

import "gorm.io/gorm"

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Nome     string `json:"name"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Nome  string `json:"nome"`
	Email string `json:"email" gorm:"unique"`
	Senha string `json:"senha"`
	Admin bool   `json:"admin"`
}

type LoginResponse struct {
	Token string `json:"token"`
	CommonResponse
}
type RegisterResponse LoginResponse

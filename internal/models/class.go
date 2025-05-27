package models

import "time"

type Class struct {
	Id        int64      `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Name      string     `json:"name"`
}

type ClassRequest struct {
	Classes []Class `json:"classes"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ForgotPasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

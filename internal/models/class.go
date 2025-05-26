package models

import "time"

type Class struct {
	Id        int64      `json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Name      string     `json:"name"`
}

type ClassRequest struct {
	Classes []Class `json:"classes"`
}

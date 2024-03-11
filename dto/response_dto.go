package dto

import "time"

type ResponseDTO struct {
	ID        uint       `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type BaseResponse struct {
	Message string      `json:"message"`
	Errors  string      `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
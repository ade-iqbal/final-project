package dto

type UserRequest struct {
	Age      uint   `json:"age,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}

type UserResponse struct {
	ResponseDTO
	Age       uint       `json:"age,omitempty"`
	Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	Username  string     `json:"username,omitempty"`
}
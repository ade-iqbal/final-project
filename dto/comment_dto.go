package dto

type CommentRequest struct {
	PhotoID uint   `json:"photo_id,omitempty"`
	Message string `json:"message,omitempty"`
}

type CommentResponse struct {
	ResponseDTO
	UserID  uint   `json:"user_id,omitempty"`
	PhotoID uint   `json:"photo_id,omitempty"`
	Message string `json:"message,omitempty"`

	User  *UserResponse  `json:"user,omitempty"`
	Photo *PhotoResponse `json:"photo,omitempty"`
}
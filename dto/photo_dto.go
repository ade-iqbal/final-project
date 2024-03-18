package dto

type PhotoRequest struct {
	Title    string `json:"title,omitempty"`
	Caption  string `json:"caption,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	UserID   uint   `json:"user_id,omitempty"`
}

type PhotoResponse struct {
	ResponseDTO
	Title    string `json:"title,omitempty"`
	Caption  string `json:"caption,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	UserID   uint   `json:"user_id,omitempty"`

	User *UserResponse `json:"user,omitempty"`
}
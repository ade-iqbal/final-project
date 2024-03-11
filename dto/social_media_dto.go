package dto

type SocialMediaRequest struct {
	Name           string `json:"name,omitempty"`
	SocialMediaUrl string `json:"social_media_url,omitempty"`
}

type SocialMediaResponse struct {
	ResponseDTO
	Name           string `json:"name,omitempty"`
	SocialMediaUrl string `json:"social_media_url,omitempty"`
	UserID         uint   `json:"user_id,omitempty"`

	User *UserResponse `json:",omitempty"`
}
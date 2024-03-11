package repository

import (
	"fga-final-project/model"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	GetAllSocialMedia() (*[]model.SocialMedia, error)
	GetSocialMediaById(id *uint) (*model.SocialMedia, error)
	GetSocialMediaByIdAndUserId(id *uint, userId *uint) (*model.SocialMedia, error)
	CreateSocialMedia(socialMedia *model.SocialMedia) error
	UpdateSocialMedia(socialMedia *model.SocialMedia) error
	DeleteSocialMedia(id *uint) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) GetAllSocialMedia() (*[]model.SocialMedia, error) {
	var socialMedias []model.SocialMedia
	err := r.db.Preload("User").Find(&socialMedias).Error
	return &socialMedias, err
}

func (r *socialMediaRepository) GetSocialMediaById(id *uint) (*model.SocialMedia, error) {
	var socialMedia model.SocialMedia
	err := r.db.Where("id = ?", *id).Take(&socialMedia).Error
	return &socialMedia, err
}

func (r *socialMediaRepository) GetSocialMediaByIdAndUserId(id *uint, userId *uint) (*model.SocialMedia, error) {
	var socialMedia model.SocialMedia
	err := r.db.Where("id = ? AND user_id = ?", *id, *userId).Take(&socialMedia).Error
	return &socialMedia, err
}

func (r *socialMediaRepository) CreateSocialMedia(socialMedia *model.SocialMedia) error {
	return r.db.Create(socialMedia).Error
}

func (r *socialMediaRepository) UpdateSocialMedia(socialMedia *model.SocialMedia) error {
	return r.db.Save(socialMedia).Error
}

func (r *socialMediaRepository) DeleteSocialMedia(id *uint) error {
	return r.db.Where("id = ?", *id).Delete(&model.SocialMedia{}).Error
}
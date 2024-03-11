package model

import (
	"fga-final-project/helper"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" gorm:"not null" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" validate:"required"`
	UserID         uint   `json:"user_id"`
	
	User           *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (sm *SocialMedia) BeforeCreate(trx *gorm.DB) (err error) {
	validate := validator.New()

	if errCreate := validate.Struct(*sm); errCreate != nil {
		helper.LoggingError("Validation Error", err)
		err = errCreate
		return
	}

	err = nil
	return
}

func (sm *SocialMedia) BeforeUpdate(trx *gorm.DB) (err error) {
	validate := validator.New()

	if errUpdate := validate.Struct(*sm); errUpdate != nil {
		helper.LoggingError("Validation Error", err)
		err = errUpdate
		return
	}

	err = nil
	return
}
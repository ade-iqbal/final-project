package model

import (
	"fga-final-project/helper"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message" gorm:"not null" validate:"required"`

	User    *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photo   *Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (comment *Comment) BeforeCreate(trx *gorm.DB) (err error) {
	validate := validator.New()

	if errCreate := validate.Struct(*comment); errCreate != nil {
		helper.LoggingError("Validation Error", err)
		err = errCreate
		return
	}

	err = nil
	return
}

func (comment *Comment) BeforeUpdate(trx *gorm.DB) (err error) {
	validate := validator.New()

	if errUpdate := validate.Struct(*comment); errUpdate != nil {
		helper.LoggingError("Validation Error", err)
		err = errUpdate
		return
	}

	err = nil
	return
}
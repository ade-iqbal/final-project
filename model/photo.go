package model

import (
	"fga-final-project/helper"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" validate:"required"`
	Caption  *string 
	PhotoUrl string `gorm:"not null" validate:"required"`
	UserID   uint   
	
	User     *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (photo *Photo) BeforeCreate(trx *gorm.DB) (err error) {
	validate := validator.New()

	if errCreate := validate.Struct(*photo); errCreate != nil {
		helper.LoggingError("Validation Error", err)
		err = errCreate
		return
	}

	err = nil
	return
}

func (photo *Photo) BeforeUpdate(trx *gorm.DB) (err error) {
	validate := validator.New()

	if errUpdate := validate.Struct(*photo); errUpdate != nil {
		helper.LoggingError("Validation Error", err)
		err = errUpdate
		return
	}

	err = nil
	return
}
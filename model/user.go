package model

import (
	"fga-final-project/helper"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" validate:"required"`
	Email    string `gorm:"not null;uniqueIndex" validate:"required,email"`
	Password string `gorm:"not null" validate:"required,min=6"`
	Age      uint   `gorm:"not null" validate:"required,gt=8"`
}

func (user *User) BeforeCreate(trx *gorm.DB) (err error) {
	validate := validator.New()

	if errCreate := validate.Struct(*user); errCreate != nil {
		helper.LoggingError("Validation Error", err)
		err = errCreate
		return
	}

	user.Password = helper.HassPassword(user.Password)
	err = nil
	return
}

func (user *User) BeforeUpdate(trx *gorm.DB) (err error) {
	validate := validator.New()

	if errUpdate := validate.Struct(*user); errUpdate != nil {
		helper.LoggingError("Validation Error", err)
		err = errUpdate
		return
	}

	err = nil
	return
}
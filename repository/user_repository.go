package repository

import (
	"fga-final-project/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(id *uint) (*model.User, error)
	GetUserByEmail(email *string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id *uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserById(id *uint) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", *id).Take(&user).Error
	return &user, err
}

func (r *userRepository) GetUserByEmail(email *string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", *email).Take(&user).Error
	return &user, err
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id *uint) (error) {
	return r.db.Where("id = ?", *id).Delete(&model.User{}).Error
}
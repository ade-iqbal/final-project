package repository

import (
	"fga-final-project/model"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	GetAllPhoto() (*[]model.Photo, error)
	GetPhotoById(id *uint) (*model.Photo, error)
	GetPhotoByIdAndUserId(id *uint, userId *uint) (*model.Photo, error)
	CreatePhoto(photo *model.Photo) error
	UpdatePhoto(photo *model.Photo) error
	DeletePhoto(id *uint) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) GetAllPhoto() (*[]model.Photo, error) {
	var photos []model.Photo
	err := r.db.Preload("User").Find(&photos).Error
	return &photos, err
}

func (r *photoRepository) GetPhotoById(id *uint) (*model.Photo, error) {
	var photo model.Photo
	err := r.db.Where("id = ?", *id).Take(&photo).Error
	return &photo, err
}

func (r *photoRepository) GetPhotoByIdAndUserId(id *uint, userId *uint) (*model.Photo, error) {
	var photo model.Photo
	err := r.db.Where("id = ? AND user_id = ?", *id, *userId).Take(&photo).Error
	return &photo, err
}

func (r *photoRepository) CreatePhoto(photo *model.Photo) error {
	return r.db.Create(photo).Error
}

func (r *photoRepository) UpdatePhoto(photo *model.Photo) error {
	return r.db.Save(photo).Error
}

func (r *photoRepository) DeletePhoto(id *uint) error {
	return r.db.Where("id = ?", *id).Delete(&model.Photo{}).Error
}
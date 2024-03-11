package service

import (
	"fga-final-project/helper"
	"fga-final-project/model"
	"fga-final-project/repository"
)

type PhotoService interface {
	CreatePhoto(photo model.Photo) (model.Photo, error)
	GetAllPhoto() ([]model.Photo, error)
	UpdatePhoto(id uint, photo model.Photo) (model.Photo, error)
	DeletePhoto(id uint) error
}

type photoService struct {
	PhotoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository) PhotoService {
	return &photoService{photoRepository}
}

func (s *photoService) CreatePhoto(photo model.Photo) (model.Photo, error) {
	err := s.PhotoRepository.CreatePhoto(&photo)
	if err != nil {
		helper.LoggingError("Create Data Error", err)
	}

	return photo, err
}

func (s *photoService) GetAllPhoto() ([]model.Photo, error) {
	photos, err := s.PhotoRepository.GetAllPhoto()
	if err != nil {
		helper.LoggingError("Data Not Found", err)
	}

	return *photos, err
}

func (s *photoService) UpdatePhoto(id uint, photo model.Photo) (model.Photo, error) {
	var (
		photoData *model.Photo
		err error
	)

	photoData, err = s.PhotoRepository.GetPhotoById(&id)
	if err != nil {
		helper.LoggingError("Data Not Found", err)
		return *photoData, err
	}

	photoData.Title = photo.Title
	photoData.Caption = photo.Caption
	photoData.PhotoUrl = photo.PhotoUrl
	
	err = s.PhotoRepository.UpdatePhoto(photoData)
	if err != nil {
		helper.LoggingError("Update Data Error", err)
	}

	return *photoData, err
}

func (s *photoService) DeletePhoto(id uint) error {
	err := s.PhotoRepository.DeletePhoto(&id)
	if err != nil {
		helper.LoggingError("Delete Data Error", err)
	}

	return err
}
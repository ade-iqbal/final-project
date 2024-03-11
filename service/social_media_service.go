package service

import (
	"fga-final-project/helper"
	"fga-final-project/model"
	"fga-final-project/repository"
)

type SocialMediaService interface {
	CreateSocialMedia(socialMedia model.SocialMedia) (model.SocialMedia, error)
	GetAllSocialMedia() ([]model.SocialMedia, error)
	UpdateSocialMedia(id uint, socialMedia model.SocialMedia) (model.SocialMedia, error)
	DeleteSocialMedia(id uint) error
}

type socialMediaService struct {
	SocialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{socialMediaRepository}
}

func (s *socialMediaService) CreateSocialMedia(socialMedia model.SocialMedia) (model.SocialMedia, error) {
	err := s.SocialMediaRepository.CreateSocialMedia(&socialMedia)
	if err != nil {
		helper.LoggingError("Create Data Error", err)
	}

	return socialMedia, err
}

func (s *socialMediaService) GetAllSocialMedia() ([]model.SocialMedia, error) {
	socialMedias, err := s.SocialMediaRepository.GetAllSocialMedia()
	if err != nil {
		helper.LoggingError("Data Not Found", err)
	}

	return *socialMedias, err
}

func (s *socialMediaService) UpdateSocialMedia(id uint, socialMedia model.SocialMedia) (model.SocialMedia, error) {
	var (
		socialMediaData *model.SocialMedia
		err error
	)

	socialMediaData, err = s.SocialMediaRepository.GetSocialMediaById(&id)
	if err != nil {
		helper.LoggingError("Data Not Found", err)
		return *socialMediaData, err
	}

	socialMediaData.Name = socialMedia.Name
	socialMediaData.SocialMediaUrl = socialMedia.SocialMediaUrl

	err = s.SocialMediaRepository.UpdateSocialMedia(socialMediaData)
	if err != nil {
		helper.LoggingError("Update Data Error", err)
	}

	return *socialMediaData, err
}

func (s *socialMediaService) DeleteSocialMedia(id uint) error {
	err := s.SocialMediaRepository.DeleteSocialMedia(&id)
	if err != nil {
		helper.LoggingError("Delete Data Error", err)
	}

	return err
}

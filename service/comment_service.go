package service

import (
	"fga-final-project/helper"
	"fga-final-project/model"
	"fga-final-project/repository"
)

type CommentService interface {
	CreateComment(comment model.Comment) (model.Comment, error)
	GetAllComment() ([]model.Comment, error)
	UpdateComment(id uint, comment model.Comment) (model.Comment, error)
	DeleteComment(id uint) error
}

type commentService struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) CommentService {
	return &commentService{commentRepository}
}

func (s *commentService) CreateComment(comment model.Comment) (model.Comment, error) {
	err := s.CommentRepository.CreateComment(&comment)
	if err != nil {
		helper.LoggingError("Create Data Error", err)
	}

	return comment, err
}

func (s *commentService) GetAllComment() ([]model.Comment, error) {
	comments, err := s.CommentRepository.GetAllComment()
	if err != nil {
		helper.LoggingError("Data Not Found", err)
	}

	return *comments, err
}

func (s *commentService) UpdateComment(id uint, comment model.Comment) (model.Comment, error) {
	var (
		commentData *model.Comment
		err error
	)

	commentData, err = s.CommentRepository.GetCommentById(&id)
	if err != nil {
		helper.LoggingError("Data Not Found", err)
		return *commentData, err
	}

	commentData.Message = comment.Message

	err = s.CommentRepository.UpdateComment(commentData)
	if err != nil {
		helper.LoggingError("Update Data Error", err)
	}

	return *commentData, err
}

func (s *commentService) DeleteComment(id uint) error {
	err := s.CommentRepository.DeleteComment(&id)
	if err != nil {
		helper.LoggingError("Delete Data Error", err)
	}

	return err
}
package repository

import (
	"fga-final-project/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	GetAllComment() (*[]model.Comment, error)
	GetCommentById(id *uint) (*model.Comment, error)
	GetCommentByIdAndUserId(id *uint, userId *uint) (*model.Comment, error)
	CreateComment(comment *model.Comment) error
	UpdateComment(comment *model.Comment) error
	DeleteComment(id *uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) GetAllComment() (*[]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	return &comments, err
}

func (r *commentRepository) GetCommentById(id *uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.Where("id = ?", *id).Take(&comment).Error
	return &comment, err
}

func (r *commentRepository) GetCommentByIdAndUserId(id *uint, userId *uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.Where("id = ? AND user_id = ?", *id, *userId).Take(&comment).Error
	return &comment, err
}

func (r *commentRepository) CreateComment(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) UpdateComment(comment *model.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) DeleteComment(id *uint) error {
	return r.db.Where("id = ?", *id).Delete(&model.Comment{}).Error
}
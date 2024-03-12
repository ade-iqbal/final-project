package handler

import (
	"errors"
	"fga-final-project/dto"
	"fga-final-project/model"
	"fga-final-project/service"
	"fga-final-project/util"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentHandler interface {
	CreateComment(ctx *gin.Context)
	GetAllComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentHandler struct {
	CommentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) CommentHandler {
	return &commentHandler{commentService}
}

func (h *commentHandler) CreateComment(ctx *gin.Context) {
	var (
		commentRequest dto.CommentRequest

		userData = ctx.MustGet("UserData").(jwt.MapClaims)
		userId = uint(userData["id"].(float64))
	)

	ctx.ShouldBindJSON(&commentRequest)
	comment := model.Comment {
		Message: commentRequest.Message,
		PhotoID: commentRequest.PhotoID,
		UserID: userId,
	}

	newComment, err := h.CommentService.CreateComment(comment)
	if err != nil {
		baseResponse := dto.BaseResponse {
			Message: util.InvalidDataMessage,
			Errors: err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}

	commentResponse := dto.CommentResponse {
		ResponseDTO: dto.ResponseDTO {
			ID: newComment.ID,
			CreatedAt: newComment.CreatedAt,
		},
		Message: newComment.Message,
		PhotoID: newComment.PhotoID,
		UserID: newComment.UserID,
	}

	ctx.JSON(http.StatusCreated, commentResponse)
}

func (h *commentHandler) GetAllComment(ctx *gin.Context) {
	var (
		commentsResponse []interface{}
		commentResponse dto.CommentResponse
	)

	comments, err := h.CommentService.GetAllComment()
	if err != nil {
		baseResponse := dto.BaseResponse {
			Message: util.DataNotFoundMessage,
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, baseResponse)
		return
	}

	for _, value := range comments {
		commentResponse = dto.CommentResponse {
			ResponseDTO: dto.ResponseDTO {
				ID: value.ID,
				UpdatedAt: value.UpdatedAt,
				CreatedAt: value.CreatedAt,
			},
			Message: value.Message,
			PhotoID: value.PhotoID,
			UserID: value.UserID,
			User: &dto.UserResponse {
				ResponseDTO: dto.ResponseDTO {
					ID: value.User.ID,
				},
				Email: value.User.Email,
				Username: value.User.Username,
			},
			Photo: &dto.PhotoResponse {
				ResponseDTO: dto.ResponseDTO {
					ID: value.Photo.ID,
				},
				Title: value.Photo.Title,
				Caption: *value.Photo.Caption,
				PhotoUrl: value.Photo.PhotoUrl,
				UserID: value.Photo.UserID,
			},
		}

		commentsResponse = append(commentsResponse, commentResponse)
	}

	ctx.JSON(http.StatusOK, commentsResponse)
}

func (h *commentHandler) UpdateComment(ctx *gin.Context) {
	var (
		commentRequest dto.CommentRequest

		commentId, _ = strconv.Atoi(ctx.Param("commentId"))
	)

	ctx.ShouldBindJSON(&commentRequest)
	comment := model.Comment {
		Message: commentRequest.Message,
	}

	updatedComment, err := h.CommentService.UpdateComment(uint(commentId), comment)
	if err != nil {
		var baseResponse dto.BaseResponse

		switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				baseResponse.Message = util.DataNotFoundMessage
				ctx.AbortWithStatusJSON(http.StatusNotFound, baseResponse)

			default:
				baseResponse.Message = util.InvalidDataMessage
				baseResponse.Errors = err.Error()
				ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		}

		return
	}

	// TODO: konfirmasi lagi responsenya benar atau salah
	// di deskripsi mengirimkan response photo bukan comment
	commentResponse := dto.CommentResponse {
		ResponseDTO: dto.ResponseDTO {
			ID: updatedComment.ID,
			UpdatedAt: updatedComment.UpdatedAt,
		},
		Message: updatedComment.Message,	
		UserID: updatedComment.UserID,
		PhotoID: updatedComment.PhotoID,
	}

	ctx.JSON(http.StatusOK, commentResponse)
}

func (h *commentHandler) DeleteComment(ctx *gin.Context) {
	var (
		baseResponse dto.BaseResponse
		
		commentId, _ = strconv.Atoi(ctx.Param("commentId"))
	)

	err := h.CommentService.DeleteComment(uint(commentId))
	if err != nil {
		baseResponse.Message = util.InvalidDataMessage
		baseResponse.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}

	baseResponse.Message = "Your comment has been successfully deleted"
	ctx.JSON(http.StatusOK, baseResponse)
}
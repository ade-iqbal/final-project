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

type PhotoHandler interface {
	CreatePhoto(ctx *gin.Context) 
	GetAllPhoto(ctx *gin.Context) 
	UpdatePhoto(ctx *gin.Context) 
	DeletePhoto(ctx *gin.Context) 
}

type photoHandler struct {
	PhotoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) PhotoHandler {
	return &photoHandler{photoService}
}

func (h *photoHandler) CreatePhoto(ctx *gin.Context) {
	var (
		photoRequest dto.PhotoRequest

		userData = ctx.MustGet("UserData").(jwt.MapClaims)
		userId = uint(userData["id"].(float64))
	)

	ctx.ShouldBindJSON(&photoRequest)
	photo := model.Photo {
		Title: photoRequest.Title,
		Caption: &photoRequest.Caption,
		PhotoUrl: photoRequest.PhotoUrl,
		UserID: userId,
	}

	newPhoto, err := h.PhotoService.CreatePhoto(photo)
	if err != nil {
		baseResponse := dto.BaseResponse {
			Message: util.InvalidDataMessage,
			Errors: err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}

	photoResponse := dto.PhotoResponse {
		ResponseDTO: dto.ResponseDTO {
			ID: newPhoto.ID,
			CreatedAt: newPhoto.CreatedAt,
		},
		Title: newPhoto.Title,
		Caption: *newPhoto.Caption,
		PhotoUrl: newPhoto.PhotoUrl,
		UserID: newPhoto.UserID,
	}

	ctx.JSON(http.StatusCreated, photoResponse)
}

func (h *photoHandler) GetAllPhoto(ctx *gin.Context) {
	var (
		photosResponse []interface{}
		photoResponse dto.PhotoResponse
	)

	photos, err := h.PhotoService.GetAllPhoto()
	if err != nil {
		baseResponse := dto.BaseResponse {
			Message: util.DataNotFoundMessage,
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, baseResponse)
		return
	}

	for _, value := range photos {
		photoResponse = dto.PhotoResponse{
			ResponseDTO: dto.ResponseDTO {
				ID: value.ID,
				CreatedAt: value.CreatedAt,
				UpdatedAt: value.UpdatedAt,
			},
			Title: value.Title,
			Caption: *value.Caption,
			PhotoUrl: value.PhotoUrl,
			UserID: value.UserID,
			User: &dto.UserResponse {
				Email: value.User.Email,
				Username: value.User.Username,
			},
		}

		photosResponse = append(photosResponse, photoResponse)		
	}

	ctx.JSON(http.StatusOK, photosResponse)
}

func (h *photoHandler) UpdatePhoto(ctx *gin.Context) {
	var (
		photoRequest dto.PhotoRequest

		photoId, _ = strconv.Atoi(ctx.Param("photoId"))
	)

	ctx.ShouldBindJSON(&photoRequest)
	photo := model.Photo {
		Title: photoRequest.Title,
		Caption: &photoRequest.Caption,
		PhotoUrl: photoRequest.PhotoUrl,
	}

	updatedPhoto, err := h.PhotoService.UpdatePhoto(uint(photoId), photo)
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

	photoResponse := dto.PhotoResponse {
		ResponseDTO: dto.ResponseDTO {
			ID: updatedPhoto.ID,
			UpdatedAt: updatedPhoto.UpdatedAt,
		},
		Title: updatedPhoto.Title,
		Caption: *updatedPhoto.Caption,
		PhotoUrl: updatedPhoto.PhotoUrl,
		UserID: updatedPhoto.UserID,
	}

	ctx.JSON(http.StatusOK, photoResponse)
}

func (h *photoHandler) DeletePhoto(ctx *gin.Context) {
	var (
		baseResponse dto.BaseResponse
		
		photoId, _ = strconv.Atoi(ctx.Param("photoId"))
	)

	err := h.PhotoService.DeletePhoto(uint(photoId))
	if err != nil {
		baseResponse.Message = util.InvalidDataMessage
		baseResponse.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}

	baseResponse.Message = "Your photo has been successfully deleted"
	ctx.JSON(http.StatusOK, baseResponse)
}
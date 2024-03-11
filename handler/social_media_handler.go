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

type SocialMediaHandler interface {
	CreateSocialMedia(ctx *gin.Context)
	GetAllSocialMedia(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaHandler struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) SocialMediaHandler {
	return &socialMediaHandler{socialMediaService}
}

func (h *socialMediaHandler) CreateSocialMedia(ctx *gin.Context) {
	var (
		socialMediaRequest dto.SocialMediaRequest

		userData = ctx.MustGet("UserData").(jwt.MapClaims)
		userId = uint(userData["id"].(float64))
	)

	ctx.ShouldBindJSON(&socialMediaRequest)
	socialMedia := model.SocialMedia {
		Name: socialMediaRequest.Name,
		SocialMediaUrl: socialMediaRequest.SocialMediaUrl,
		UserID: userId,
	}

	newSocialMedia, err := h.SocialMediaService.CreateSocialMedia(socialMedia)
	if err != nil {
		baseResponse := dto.BaseResponse {
			Message: util.InvalidDataMessage,
			Errors: err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}
	
	socialMediaResponse := dto.SocialMediaResponse {
		ResponseDTO: dto.ResponseDTO {
			ID: newSocialMedia.ID,
			CreatedAt: newSocialMedia.CreatedAt,
		},
		Name: newSocialMedia.Name,
		SocialMediaUrl: newSocialMedia.SocialMediaUrl,
		UserID: newSocialMedia.UserID,
	}

	ctx.JSON(http.StatusCreated, socialMediaResponse)
}

func (h *socialMediaHandler) GetAllSocialMedia(ctx *gin.Context) {
	var (
		socialMediasResponse []interface{}
		socialMediaResponse dto.SocialMediaResponse
	)

	socialMedias, err := h.SocialMediaService.GetAllSocialMedia()
	if err != nil {
		baseResponse := dto.BaseResponse {
			Message: util.DataNotFoundMessage,
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, baseResponse)
		return
	}

	// TODO: konfirmasi lagi responsenya benar atau salah
	// penamaan key di deskripsi tidak konsisten
	for _, value := range socialMedias {
		socialMediaResponse = dto.SocialMediaResponse {
			ResponseDTO: dto.ResponseDTO {
				ID: value.ID,
				CreatedAt: value.CreatedAt,
				UpdatedAt: value.UpdatedAt,
			},
			Name: value.Name,
			SocialMediaUrl: value.SocialMediaUrl,
			UserID: value.UserID,
			User: &dto.UserResponse {
				ResponseDTO: dto.ResponseDTO {
					ID: value.User.ID,
				},
				Username: value.User.Username,
				Email: value.User.Email, // di deskripsi mengirimkan data photo url
			},
		}

		socialMediasResponse = append(socialMediasResponse, socialMediaResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"social_medias": socialMediasResponse,
	})
}

func (h *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	var (
		socialMediaRequest dto.SocialMediaRequest

		socialMediaId, _ = strconv.Atoi(ctx.Param("socialMediaId"))
	)

	ctx.ShouldBindJSON(&socialMediaRequest)
	socialMedia := model.SocialMedia {
		Name: socialMediaRequest.Name,
		SocialMediaUrl: socialMediaRequest.SocialMediaUrl,
	}

	updatedSocialMedia, err := h.SocialMediaService.UpdateSocialMedia(uint(socialMediaId), socialMedia)
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

	socialMediaResponse := dto.SocialMediaResponse {
		ResponseDTO: dto.ResponseDTO {
			ID: updatedSocialMedia.ID,
			UpdatedAt: socialMedia.UpdatedAt,
		},
		Name: updatedSocialMedia.Name,
		SocialMediaUrl: updatedSocialMedia.SocialMediaUrl,
		UserID: updatedSocialMedia.UserID,
	}

	ctx.JSON(http.StatusOK, socialMediaResponse)
}

func (h *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
	var (
		baseResponse dto.BaseResponse
		
		socialMediaId, _ = strconv.Atoi(ctx.Param("socialMediaId"))
	)

	err := h.SocialMediaService.DeleteSocialMedia(uint(socialMediaId))
	if err != nil {
		baseResponse.Message = util.InvalidDataMessage
		baseResponse.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}

	baseResponse.Message = "Your social media has been successfully deleted"
	ctx.JSON(http.StatusOK, baseResponse)
}

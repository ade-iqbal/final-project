package middleware

import (
	"errors"
	"fga-final-project/config"
	"fga-final-project/dto"
	"fga-final-project/helper"
	"fga-final-project/model"
	"fga-final-project/repository"
	"fga-final-project/util"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	errUserNotFound = errors.New("user is not found")
	err error
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			db = config.GetDB()
			baseResponse dto.BaseResponse
			idUser uint

			userData = ctx.MustGet("UserData").(jwt.MapClaims)
			userId = uint(userData["id"].(float64))
			paramId, err = strconv.Atoi(ctx.Param(ctx.Params[0].Key))
		)

		if err != nil {
			helper.LoggingError("Authorization Error", err)

			baseResponse.Message = util.InvalidDataMessage
			baseResponse.Errors = err.Error()
			ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
			return
		}

		switch ctx.Params[0].Key {
			case "photoId":
				idUser, err = authPhoto(db, uint(paramId), userId)
			case "commentId":
				idUser, err = authComment(db, uint(paramId), userId)
			case "socialMediaId":
				idUser, err = authSocialMedia(db, uint(paramId), userId)
		}

		if err != nil {
			helper.LoggingError("Authorization Error", err)

			baseResponse.Message = util.ForbiddenMessage
			ctx.AbortWithStatusJSON(http.StatusNotFound, baseResponse)
			return
		}

		if idUser != userId {
			helper.LoggingError("Authorization Error", err)

			baseResponse.Message = util.ForbiddenMessage
			ctx.AbortWithStatusJSON(http.StatusForbidden, baseResponse)
			return
		}

		ctx.Next()
	}
}

func authPhoto(db *gorm.DB, paramId uint, userId uint) (uint, error) {
	var (
		photoRepository = repository.NewPhotoRepository(db)
		photo *model.Photo
	)

	if photo, err = photoRepository.GetPhotoByIdAndUserId(&paramId, &userId); err != nil {
		return 0, err
	}

	return photo.UserID, nil
}

func authComment(db *gorm.DB, paramId uint, userId uint) (uint, error) {
	var (
		commentRepository = repository.NewCommentRepository(db)
		comment *model.Comment
	)

	if comment, err = commentRepository.GetCommentByIdAndUserId(&paramId, &userId); err != nil {
		return 0, errUserNotFound
	}

	return comment.UserID, nil
}

func authSocialMedia(db *gorm.DB, paramId uint, userId uint) (uint, error) {
	var (
		socialMediaRepository = repository.NewSocialMediaRepository(db)
		socialMedia *model.SocialMedia
	)

	if socialMedia, err = socialMediaRepository.GetSocialMediaByIdAndUserId(&paramId, &userId); err != nil {
		return 0, errUserNotFound
	}

	return socialMedia.UserID, nil
}
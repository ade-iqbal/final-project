package middleware

import (
	"fga-final-project/config"
	"fga-final-project/dto"
	"fga-final-project/helper"
	"fga-final-project/repository"
	"fga-final-project/util"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	err error
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			db = config.GetDB()
			baseResponse dto.BaseResponse

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
				err = authPhoto(db, uint(paramId), userId)
			case "commentId":
				err = authComment(db, uint(paramId), userId)
			case "socialMediaId":
				err = authSocialMedia(db, uint(paramId), userId)
		}

		if err != nil {
			helper.LoggingError("Authorization Error", err)

			baseResponse.Message = util.ForbiddenMessage
			ctx.AbortWithStatusJSON(http.StatusNotFound, baseResponse)
			return
		}

		ctx.Next()
	}
}

func authPhoto(db *gorm.DB, paramId uint, userId uint) error {
	photoRepository := repository.NewPhotoRepository(db)
	_, err = photoRepository.GetPhotoByIdAndUserId(&paramId, &userId)
	return err
}

func authComment(db *gorm.DB, paramId uint, userId uint) error {
	commentRepository := repository.NewCommentRepository(db)
	_, err = commentRepository.GetCommentByIdAndUserId(&paramId, &userId)
	return  err
}

func authSocialMedia(db *gorm.DB, paramId uint, userId uint) error {
	socialMediaRepository := repository.NewSocialMediaRepository(db)
	_, err = socialMediaRepository.GetSocialMediaByIdAndUserId(&paramId, &userId)
	return err
}
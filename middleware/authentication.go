package middleware

import (
	"fga-final-project/dto"
	"fga-final-project/helper"
	"fga-final-project/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var BaseResponse dto.BaseResponse
		verifyToken, err := helper.VerifyToken(ctx)

		if err != nil {
			helper.LoggingError("Authentication Error", err)

			BaseResponse.Message = util.UnauthorizedMessage
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, BaseResponse)
			return
		}
		
		ctx.Set("UserData", verifyToken)
		ctx.Next()
	}
}
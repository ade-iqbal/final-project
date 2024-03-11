package helper

import (
	"github.com/gin-gonic/gin"
)

const (
	json = "application/json"
	formUrlEncoded = "application/x-www-form-urlencoded"
)

func GetContentType(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}

func GetToken(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Authorization")
}

func BindingData(T interface{}, ctx *gin.Context) {
	switch GetContentType(ctx) {
		case json:
			ctx.ShouldBindJSON(&T)
		case formUrlEncoded:
			ctx.ShouldBind(&T)
	}
}
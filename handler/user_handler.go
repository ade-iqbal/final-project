package handler

import (
	"errors"
	"fga-final-project/dto"
	"fga-final-project/helper"
	"fga-final-project/model"
	"fga-final-project/service"
	"fga-final-project/util"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler interface {
	Register(ctx *gin.Context) 
	Login(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(ctx *gin.Context) {
	var (
		userRequest dto.UserRequest
	)

	ctx.ShouldBindJSON(&userRequest)
	user := model.User {
		Age: userRequest.Age,
		Email: userRequest.Email,
		Password: userRequest.Password,
		Username: userRequest.Username,
	}

	newUser, err := h.UserService.Register(user)
	if err != nil {
		baseResponse := dto.BaseResponse {
			Message: util.InvalidDataMessage,
			Errors: err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}
	
	userResponse := dto.UserResponse {
		Age: newUser.Age,
		Email: newUser.Email,
		ResponseDTO: dto.ResponseDTO {
			ID: newUser.ID,
		},
		Username: newUser.Username,
	}

	ctx.JSON(http.StatusCreated, userResponse)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var (
		userRequest dto.UserRequest
	)

	ctx.ShouldBindJSON(&userRequest)
	user := model.User {
		Email: userRequest.Email,
		Password: userRequest.Password,
	}

	token, err := h.UserService.Login(user)
	if err != nil {
		baseResponse := dto.BaseResponse {
			Message: util.InvalidCredentialsMessage,
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, baseResponse)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var (
		userRequest dto.UserRequest
		baseResponse dto.BaseResponse

		userData = ctx.MustGet("UserData").(jwt.MapClaims)
		userId = uint(userData["id"].(float64))
		userIdParam, _ = strconv.Atoi(ctx.Param("userId"))
	)

	if userId != uint(userIdParam) {
		helper.LoggingError("Forbidden Access", util.ErrForbidden)

		baseResponse.Message = util.ForbiddenMessage
		ctx.AbortWithStatusJSON(http.StatusForbidden, baseResponse)
		return
	}

	ctx.ShouldBindJSON(&userRequest)
	user := model.User {
		Email: userRequest.Email,
		Username: userRequest.Username,
	}

	updatedUser, err := h.UserService.UpdateUser(uint(userIdParam), user)
	if err != nil {
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

	userResponse := dto.UserResponse {
		ResponseDTO: dto.ResponseDTO {
			ID: updatedUser.ID,
			UpdatedAt: updatedUser.UpdatedAt,
		},
		Email: updatedUser.Email,
		Username: updatedUser.Username,
		Age: updatedUser.Age,
	}

	ctx.JSON(http.StatusOK, userResponse)
}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	var (
		baseResponse dto.BaseResponse
		err error

		userData = ctx.MustGet("UserData").(jwt.MapClaims)
		userId = uint(userData["id"].(float64))
	)

	err = h.UserService.DeleteUser(userId)
	if err != nil {
		baseResponse.Message = util.InvalidDataMessage
		baseResponse.Errors = err.Error()

		ctx.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}

	baseResponse.Message = "Your account has been successfully deleted"
	ctx.JSON(http.StatusOK, baseResponse)
}
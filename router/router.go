package router

import (
	"fga-final-project/config"
	"fga-final-project/handler"
	"fga-final-project/middleware"
	"fga-final-project/repository"
	"fga-final-project/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	userHandler handler.UserHandler
	photoHandler handler.PhotoHandler
	commentHandler handler.CommentHandler
	socialMediaHandler handler.SocialMediaHandler
)

func init() {
	config.StartDB()
	db = config.GetDB()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler = handler.NewUserHandler(userService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoHandler = handler.NewPhotoHandler(photoService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentHandler = handler.NewCommentHandler(commentService)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaHandler = handler.NewSocialMediaHandler(socialMediaService)
}

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)

		userRouter.Use(middleware.Authentication())

		userRouter.PUT("/:userId", userHandler.UpdateUser)
		userRouter.DELETE("/", userHandler.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())

		photoRouter.POST("/", photoHandler.CreatePhoto)
		photoRouter.GET("/", photoHandler.GetAllPhoto)
		photoRouter.PUT("/:photoId", middleware.Authorization(), photoHandler.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.Authorization(), photoHandler.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())

		commentRouter.POST("/", commentHandler.CreateComment)
		commentRouter.GET("/", commentHandler.GetAllComment)
		commentRouter.PUT("/:commentId", middleware.Authorization(), commentHandler.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.Authorization(), commentHandler.DeleteComment)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())

		socialMediaRouter.POST("/", socialMediaHandler.CreateSocialMedia)
		socialMediaRouter.GET("/", socialMediaHandler.GetAllSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middleware.Authorization(), socialMediaHandler.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.Authorization(), socialMediaHandler.DeleteSocialMedia)
	}

	return router
}
package main

import (
	"github.com/Alifanoveliasukma/golang_apl/config"
	"github.com/Alifanoveliasukma/golang_apl/controller"
	"github.com/Alifanoveliasukma/golang_apl/middleware"
	"github.com/Alifanoveliasukma/golang_apl/repository"
	"github.com/Alifanoveliasukma/golang_apl/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	imageRepository repository.ImageRepository = repository.NewImageRepository(db)
	jwtService      service.JWTService         = service.NewJWTService()
	userService     service.UserService        = service.NewUserService(userRepository)
	imageService    service.ImageService       = service.NewImageService(imageRepository)
	authService     service.AuthService        = service.NewAuthService(userRepository)
	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	imageController controller.ImageController = controller.NewImageController(imageService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	imageRoutes := r.Group("api/images", middleware.AuthorizeJWT(jwtService))
	{
		imageRoutes.GET("/", imageController.All)
		imageRoutes.POST("/", imageController.Insert)
		imageRoutes.GET("/:id", imageController.FindByID)
		imageRoutes.PUT("/:id", imageController.Update)
		imageRoutes.DELETE("/:id", imageController.Delete)
	}

	r.Run()
}

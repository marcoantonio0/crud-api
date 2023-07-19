package route

import (
	"github.com/gin-gonic/gin"
	"github.com/marcoantonio63/crud-api/app/config"
	"github.com/marcoantonio63/crud-api/app/controller"
	"github.com/marcoantonio63/crud-api/app/repository"
	"github.com/marcoantonio63/crud-api/app/service"
)

func UserRoutes(group *gin.RouterGroup) {
	userCollection := config.GetCollection("user")
	userRepository := repository.NewUserRepository(userCollection)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	group.POST("/", userController.CreateUserHandler)
	group.GET("/", userController.ListAllUsersHandler)
	group.GET("/:id", userController.FindByIdHandler)
}

package routes

import (
	"go-technical-test-bankina/src/handler"
	"go-technical-test-bankina/src/middleware"
	"go-technical-test-bankina/src/repository"
	"go-technical-test-bankina/src/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	routeAuth := r.Group("/auth")

	routeUser := r.Group("/user")

	userRepository := repository.NewUser(db)
	userService := user.NewService(userRepository)

	// authorization
	routeUser.Use(middleware.Authorization(userService))

	userHandler := handler.NewUser(userService)
	userHandler.AuthMount(routeAuth)

	userHandler.Mount(routeUser)
}

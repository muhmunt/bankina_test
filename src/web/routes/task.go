package routes

import (
	"go-technical-test-bankina/src/handler"
	"go-technical-test-bankina/src/middleware"
	"go-technical-test-bankina/src/repository"
	task "go-technical-test-bankina/src/task"
	"go-technical-test-bankina/src/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTaskRoutes(r *gin.RouterGroup, db *gorm.DB) {
	routeTask := r.Group("/task")

	taskRepository := repository.NewTask(db)
	taskService := task.NewService(taskRepository)

	userRepository := repository.NewUser(db)
	userService := user.NewService(userRepository)

	// authorization
	routeTask.Use(middleware.Authorization(userService))

	taskHandler := handler.NewTask(taskService)
	taskHandler.Mount(routeTask)
}

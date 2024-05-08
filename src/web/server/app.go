package server

import (
	"go-technical-test-bankina/src/web/routes"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	router := gin.Default()

	port := os.Getenv("PORT")

	api := router.Group("/")

	routes.SetupUserRoutes(api, db)
	routes.SetupTaskRoutes(api, db)

	router.Run(":" + port)
}

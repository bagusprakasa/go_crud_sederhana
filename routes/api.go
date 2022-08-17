package routes

import (
	"crud_sederhana/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	api.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	api.GET("/users", controllers.FindUser)
	api.POST("/users", controllers.CreateUser)
	api.GET("/users/:id", controllers.FindUserById)
	api.PUT("/users/:id", controllers.UpdateUser)
	api.DELETE("/users/:id", controllers.DeleteUser)
	return router
}

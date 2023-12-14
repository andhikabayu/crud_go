package routes

import (
	"crud_go/controllers"
	"crud_go/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginHandler)

	r.GET("/get-post-data", middleware.AuthMiddleware(), controllers.ConsumeApi)

	r.GET("/tasks", middleware.AuthMiddleware(), controllers.FindTasks)
	r.POST("/task", middleware.AuthMiddleware(), controllers.CreateTask)
	r.GET("/task/:id", middleware.AuthMiddleware(), controllers.FindTask)
	r.PATCH("/task/:id", middleware.AuthMiddleware(), controllers.UpdateTask)
	r.DELETE("task/:id", middleware.AuthMiddleware(), controllers.DeleteTask)
	return r
}

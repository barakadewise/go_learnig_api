package main

import (
	"net/http"
	"sample-api/server"
	"sample-api/src/controllers"
	"sample-api/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	//start the server
	s := server.StartServer()
	prefix := s.Group("/api")
	//health check route
	prefix.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	//welcom message route
	prefix.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hi welcome this is my first api while learning go!",
		})
	})

	//initilize user service and controller
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	prefix.POST("/users", userController.SaveUser)
	prefix.GET("/users", userController.GetAllUsers)
	prefix.GET("/users/:id", userController.GetUserById)
	prefix.PUT("/users/:id", userController.UpdateUser)
	prefix.DELETE("/users/:id", userController.DeleteUser)

	//start the server on port 8080
	s.Run(":8080")
}

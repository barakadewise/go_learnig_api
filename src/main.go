package src

import (
	"net/http"
	"sample-api/server"

	"github.com/gin-gonic/gin"
)

func main() {
	//start the server
	s := server.StartServer()

	//health check route
	s.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	//welcom message routeg
	s.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hi welcome this is my first api while learning go!\n Keep following unill the end thanks",
		})
	})

	//start the server on port 8080
	s.Run(":8080")
}

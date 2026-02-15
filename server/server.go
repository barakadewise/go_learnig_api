package server

import "github.com/gin-gonic/gin"

// server initilization
func StartServer() *gin.Engine {
	server := gin.Default()

	//add recovery middleware to handle panics and return a 500 error
	server.Use(gin.Recovery(), gin.Logger())
	return server

}

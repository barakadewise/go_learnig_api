package controllers

import (
	"sample-api/src/entities"
	"sample-api/src/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	services.UserServiceInterface
}

func (c *UserController) SaveUser(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	savedUser, err := c.UserServiceInterface.SaveUser(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, savedUser)
}

// func (c *UserController) GetUserById(ctx *gin.Context) (string, error) {
// 	return "", nil
// }

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, _ := c.UserServiceInterface.GetAllUsers()
	ctx.JSON(200, gin.H{"users": users})
}

func (c *UserController) updateUser() error {
	return nil
}

func (c *UserController) deleteUser() error {
	return nil
}

package controllers

import (
	"sample-api/src/entities"
	"sample-api/src/services"
	"strconv"

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

func (c *UserController) GetUserById(ctx *gin.Context) {
	var id int
	if err := ctx.ShouldBind(&id); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := c.UserServiceInterface.GetUserById(id)
	if err != nil {
		ctx.JSON(404, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "seccessfully", "data": user})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, _ := c.UserServiceInterface.GetAllUsers()
	ctx.JSON(200, gin.H{"message": "successfully", "data": users})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var user entities.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = c.UserServiceInterface.UpdateUser(id, user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "user updated successfully"})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Id providedS"})
		return
	}
	err = c.UserServiceInterface.DeleteUser(id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{
		UserServiceInterface: userService,
	}
}

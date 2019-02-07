package controller

import (
	"cuso-blog/model"
	"cuso-blog/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Controller
	Service *service.UserSerivce
}

func (u *UserController) selectAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func selectAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (u *UserController) selectByID(c *gin.Context) {
	t := c.Param("id")
	id, err := strconv.Atoi(t)

	user, err := u.Service.GetUserByID(model.NewPrimary(uint(id)))
	// user, err := u.Service.GetUserByEmail("jinzhu@gmail.com")

	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
}

func NewUserController(router *gin.RouterGroup, service *service.UserSerivce) *UserController {
	instance := new(UserController)

	instance.SetBaseURL("/users")

	instance.Router = router
	instance.Service = service

	instance.Router.GET(instance.createURL("/"), instance.selectAll)
	instance.Router.GET(instance.createURL("/:id"), instance.selectByID)

	return instance
}

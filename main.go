package main

import (
	"cuso-blog/controller"
	"cuso-blog/model"
	"cuso-blog/repository"
	"cuso-blog/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
)

type Server struct {
	gin            *gin.Engine
	userController *controller.UserController
}

func (s *Server) Run() {
	s.gin.Run()
}

func NewServer(g *gin.Engine, userController *controller.UserController) *Server {
	return &Server{
		gin:            g,
		userController: userController,
	}
}

func InitalizeDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:1234@/cuso?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("err: " + err.Error())
		return
	}
	db.AutoMigrate(&model.User{}, &model.UserInfo{})

	return
}

func BuildContainer() *dig.Container {
	container := dig.New()
	engine := gin.Default()

	container.Provide(InitalizeDB)
	container.Provide(repository.NewUserRepository)
	container.Provide(service.NewUserService)
	container.Provide(controller.NewUserController)
	container.Provide(func() *gin.Engine {
		return engine
	})
	container.Provide(func(engine *gin.Engine) *gin.RouterGroup {
		return &engine.RouterGroup
	})
	container.Provide(NewServer)

	return container
}

func main() {
	container := BuildContainer()

	err := container.Invoke(func(server *Server) {
		server.Run() // listen and serve on 0.0.0.0:8080
	})

	if err != nil {
		log.Printf(err.Error())
	}
}

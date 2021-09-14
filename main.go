package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goGoGo/controllers"
	"goGoGo/middlewares"
	"os"
)

func main() {
	fmt.Println(os.Args)
	r := gin.Default()
	ctl := controllers.New()

	r.POST("/login", ctl.Login)
	r.POST("/users", ctl.CreateUser)

	authJWT := r.Group("/", middlewares.AuthJWT)
	{

		authJWT.GET("/users", ctl.GetUsers)
		authJWT.GET("/users/:id", ctl.GetUser)
		authJWT.PUT("/users/:id", ctl.UpdateUser)
		authJWT.DELETE("/users/:id", ctl.DeleteUser)

		authJWT.POST("/friend-applies", ctl.CreateFriendApply)
	}

	r.Run()
}

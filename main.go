package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goGoGo/controllers"
	"os"
)

func main() {
	fmt.Println(os.Args)
	r := gin.Default()
	ctl := controllers.New()

	r.POST("/users", ctl.CreateUser)
	r.GET("/users", ctl.GetUsers)
	r.GET("/users/:id", ctl.GetUser)
	r.PUT("/users/:id", ctl.UpdateUser)
	r.DELETE("/users/:id", ctl.DeleteUser)

	r.POST("/friend-applies", ctl.CreateFriendApply)
	r.Run()
}

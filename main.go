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
	userRepo := controllers.New()

	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)
	r.Run()
}

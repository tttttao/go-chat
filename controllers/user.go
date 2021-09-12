package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goGoGo/models"
	"goGoGo/utils"
	"gorm.io/gorm"
	"net/http"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := utils.InitDb()
	err := db.AutoMigrate(&models.User{}, &models.FriendApply{})
	if err != nil {
		fmt.Printf("AutoMigrate failed : error=%v\n", err)
	}
	return &UserRepo{Db: db}
}

// CreateUser create user
func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUsers get users
func (repository *UserRepo) GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUser get user by id
func (repository *UserRepo) GetUser(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var user models.User
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser update user
func (repository *UserRepo) UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&user)
	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser delete user
func (repository *UserRepo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	err := models.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

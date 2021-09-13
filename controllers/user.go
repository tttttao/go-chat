package controllers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goGoGo/models"
	"goGoGo/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type Repo struct {
	Db *gorm.DB
}

func New() *Repo {
	db := utils.InitDb()
	err := db.AutoMigrate(&models.User{}, &models.FriendApply{})
	if err != nil {
		fmt.Printf("AutoMigrate failed : error=%v\n", err)
	}
	return &Repo{Db: db}
}

type Login struct {
	User     string `json:"user" binding:"required"`
	Password []byte `json:"password" binding:"required"`
}

const SecretKey = "secretKey"

func (repository *Repo) Login(c *gin.Context) {
	var login Login
	if err := c.Bind(&login); err != nil {
		panic(err.Error())
	}
	user := models.User{Name: login.User}
	models.GetUser(repository.Db, &user)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "User not exists.",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, login.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Incorrect password.",
		})
		return
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.ID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could Not Login",
		})
		panic(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// CreateUser create user
func (repository *Repo) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	user.Password, _ = bcrypt.GenerateFromPassword(user.Password, 10)
	err := models.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUsers get users
func (repository *Repo) GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUser get user by id
func (repository *Repo) GetUser(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	idInt, _ := strconv.Atoi(idStr)
	user := models.User{ID: idInt}
	err := models.GetUser(repository.Db, &user)
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
func (repository *Repo) UpdateUser(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	idInt, _ := strconv.Atoi(idStr)
	user := models.User{ID: idInt}
	err := models.GetUser(repository.Db, &user)
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
func (repository *Repo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	err := models.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

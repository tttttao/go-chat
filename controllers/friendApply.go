package controllers

import (
	"github.com/gin-gonic/gin"
	"goGoGo/models"
	"net/http"
)

// CreateFriendApply create friend apply
func (repository *Repo) CreateFriendApply(c *gin.Context) {
	var apply models.FriendApply
	c.BindJSON(&apply)
	err := models.CreateFriendApply(repository.Db, &apply)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, apply)
}

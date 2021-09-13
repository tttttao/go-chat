package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
	apiRes := models.ApiFriendApply{}
	if err := copier.Copy(&apiRes, &apply); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": apiRes,
	})
}

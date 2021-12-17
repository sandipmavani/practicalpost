package handler

import (
	"log"
	"net/http"
	"practicalpost/controller"
	"practicalpost/models"

	"github.com/gin-gonic/gin"
)

func HandleEditUser(c *gin.Context) {
	var newUser models.Users
	userId := c.Param("userId")

	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	updateCount, err := controller.EditUser(userId, &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updateCount": updateCount})
}

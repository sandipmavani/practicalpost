package controller

import (
	"log"
	"net/http"
	"practicalpost/models"

	"github.com/gin-gonic/gin"
)

func HandleCreateUser(c *gin.Context) {
	var newUser models.Users
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := createUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

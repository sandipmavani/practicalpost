package handler

import (
	"net/http"
	"practicalpost/controller"

	"github.com/gin-gonic/gin"
)

func HandleListPost(c *gin.Context) {
	userList, err := controller.GetAllPost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"list": userList})
}

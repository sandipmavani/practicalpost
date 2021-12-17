package handler

import (
	"net/http"
	"practicalpost/controller"

	"github.com/gin-gonic/gin"
)

func HandleDeletePost(c *gin.Context) {
	postId := c.Param("postId")

	deleteCount, err := controller.DeletePost(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleteCount": deleteCount})
}

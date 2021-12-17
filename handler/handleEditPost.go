package handler

import (
	"log"
	"net/http"
	"practicalpost/controller"
	"practicalpost/models"

	"github.com/gin-gonic/gin"
)

func HandleEditPost(c *gin.Context) {
	var editPost models.Posts
	postId := c.Param("postId")

	if err := c.ShouldBindJSON(&editPost); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	updateCount, err := controller.EditPost(postId, &editPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updateCount": updateCount})
}

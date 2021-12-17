package handler

import (
	"fmt"
	"log"
	"net/http"
	"practicalpost/controller"
	"practicalpost/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleCreatePost(c *gin.Context) {
	var newPost models.Posts
	fmt.Println("create Post")
	userId := fmt.Sprintf("%v", c.MustGet("userId"))
	if err := c.ShouldBindJSON(&newPost); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	newPost.Author, _ = primitive.ObjectIDFromHex(userId) //get user id from the token and add as author
	newPost.PostedOn = time.Now()
	id, err := controller.CreatePost(&newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

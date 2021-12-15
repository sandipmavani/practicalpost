package controller

import (
	"log"
	"net/http"
	"practicalpost/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(c *gin.Context) {
	var user models.Users
	username := c.Param("username")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	err := client.Database("post_db").Collection("users").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Printf("Coun't get the Post")
		c.JSON(http.StatusNotFound, gin.H{"msg": "No User Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": user.Id, "username": user.UserName})
	return
}

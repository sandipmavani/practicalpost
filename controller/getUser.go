package controller

import (
	"log"
	"net/http"
	"practicalpost/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(c *gin.Context) {
	var user models.Users
	username := c.Param("username")
	client, ctx, cancel := GetConnection()
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
func GetAllUser(c *gin.Context) {
	var userList []models.Users

	client, ctx, cancel := GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	projection := bson.D{{"password", false}}

	cursor, err := client.Database("post_db").Collection("users").Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		log.Printf("Coun't get the Post")
		c.JSON(http.StatusNotFound, gin.H{"msg": "No User Found"})
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &userList); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": userList})
	return
}

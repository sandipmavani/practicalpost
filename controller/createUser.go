package controller

import (
	"log"
	"practicalpost/models"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"golang.org/x/crypto/bcrypt"
)

func createUser(user *models.Users) (string, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	password := []byte(user.Password)
	newPassword, _ := bcrypt.GenerateFromPassword(password, 10)
	user.Password = string(newPassword)

	user.Id = primitive.NewObjectID()

	result, err := client.Database("post_db").Collection("users").InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
		log.Printf("Couldn't create the user")
	}
	uid := result.InsertedID.(primitive.ObjectID).Hex()
	return uid, err
}

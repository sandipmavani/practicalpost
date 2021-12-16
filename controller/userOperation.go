package controller

import (
	"log"
	"practicalpost/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"golang.org/x/crypto/bcrypt"
)

func createUser(user *models.Users) (string, error) {
	client, ctx, cancel := GetConnection()
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
func editUser(userId string, user *models.Users) (int64, error) {
	client, ctx, cancel := GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	password := []byte(user.Password)
	newPassword, _ := bcrypt.GenerateFromPassword(password, 10)
	user.Password = string(newPassword)
	id, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": id}
	update := bson.D{{"$set", bson.D{
		{"email", user.Email},
		{"password", user.Password},
		{"phone", user.Phone},
		{"name", user.Name},
		{"dob", user.DOB},
	}}}

	result, err := client.Database("post_db").Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
		log.Printf("Couldn't create the user")
	}
	updateCount := result.ModifiedCount
	return updateCount, err
}

package authentication

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"practicalpost/controller"
	"practicalpost/models"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	user          bool
	authenticated bool
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//login API
func LoginController(c *gin.Context) {
	var loginRequest LoginRequest
	var userResponse models.Users
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	loginResponse := LoginResponse{}
	client, ctx, cancel := controller.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	fmt.Println(loginRequest)

	cur, err := client.Database("post_db").Collection("users").Find(ctx, bson.M{"email": loginRequest.Email})

	if err != nil {

	}
	for cur.Next(context.TODO()) {

		err = cur.Decode(&userResponse)
		if err != nil {
			log.Fatal(err)
		}
	}

	loginResponse = verifyUserHash(userResponse, loginRequest)
	if loginResponse.authenticated && loginResponse.user {
		c.JSON(http.StatusOK, gin.H{
			"token": GenerateToken(userResponse.Email, userResponse.Id.Hex(), true),
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Email or Password",
		})
	}
}

// here we check email is exsist and compare hash password with db password
func verifyUserHash(response models.Users, user LoginRequest) LoginResponse {
	if len(strings.TrimSpace(response.Email)) == 0 {

		return LoginResponse{user: false, authenticated: false}
	} else {
		// Comparing the password with the hash
		err := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(user.Password))
		if err == nil {
			return LoginResponse{user: true, authenticated: true}
		}
	}
	return LoginResponse{user: true, authenticated: false}
}

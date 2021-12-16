package main

import (
	"practicalpost/controller"
	"practicalpost/controller/authentication"
	"practicalpost/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//user route
	router.POST("/users", controller.HandleCreateUser)                     //create user
	router.PUT("/users/:userId", controller.HandleEditUser)                // edit user by Id
	router.GET("/users", middleware.AuthorizeJWT(), controller.GetAllUser) //get all user list
	router.GET("/users/:username", controller.GetUser)                     // search user by username

	//auth route
	router.POST("/login", authentication.LoginController) // login user

	router.Run()
}

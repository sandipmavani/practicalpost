package main

import (
	"practicalpost/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/users", controller.HandleCreateUser)      //create user
	router.PUT("/users/:userId", controller.HandleEditUser) // edit user by Id
	router.GET("/users", controller.GetAllUser)             //get all user list
	router.GET("/users/:username", controller.GetUser)      // search user by username

	router.Run()
}

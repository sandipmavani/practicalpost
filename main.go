package main

import (
	"practicalpost/controller"
	"practicalpost/controller/authentication"
	"practicalpost/handler"
	"practicalpost/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//user route
	router.POST("/users", handler.HandleCreateUser)                        //create user
	router.PUT("/users/:userId", handler.HandleEditUser)                   // edit user by Id
	router.GET("/users", middleware.AuthorizeJWT(), controller.GetAllUser) //get all user list
	router.GET("/users/:username", controller.GetUser)                     // search user by username

	//auth route
	router.POST("/login", authentication.LoginController) // login user

	//post route
	router.POST("/posts", middleware.AuthorizeJWT(), handler.HandleCreatePost)           //create new post
	router.PUT("/posts/:postId", middleware.AuthorizeJWT(), handler.HandleEditPost)      //edit post by Id
	router.POST("/posts/list", handler.HandleListPost)                                   //get all post all author list with pagination
	router.DELETE("/posts/:postId", middleware.AuthorizeJWT(), handler.HandleDeletePost) //delete post by Id

	router.Run()
}

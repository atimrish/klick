package router

import (
	"backend/controllers/ChatController"
	"backend/controllers/FriendController"
	"backend/controllers/PostController"
	"backend/controllers/UserController"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.GET("/userinfo/:user_id", UserController.GetUserInfo)

	router.POST("/register", UserController.Register)
	router.POST("/login", UserController.Login)
	router.POST("/refresh", UserController.RefreshToken)

	router.GET("/friends/:user_id", FriendController.UserFriends)
	router.POST("/invite", FriendController.Invite)
	router.PUT("/accept", FriendController.Accept)
	router.PUT("/decline", FriendController.Decline)

	router.GET("/post", PostController.GetPosts)
	router.GET("/post/:post_id", PostController.GetPostById)
	router.POST("/post", PostController.AddPost)
	router.DELETE("/post/:post_id", PostController.DeleteById)

	router.GET("/ws", ChatController.CreateChat)
}

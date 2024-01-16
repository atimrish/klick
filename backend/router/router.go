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
	router.GET("/userinfo", UserController.GetInfoMyself)

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

	router.GET("/userchat/:user_id", ChatController.GetChatsByUserId)
	router.GET("/chat/:chat_id", ChatController.GetChatById)
	router.POST("/chat", ChatController.CreateChat)
	router.POST("/chat/:chat_id", ChatController.PushMessage)
	router.PUT("/chat/:chat_id/:message_id", ChatController.UpdateMessage)
	router.DELETE("/chat/:chat_id/:message_id", ChatController.DeleteMessage)
}

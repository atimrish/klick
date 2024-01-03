package router

import (
	"backend/controllers/FriendController"
	"backend/controllers/UserController"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.POST("/register", UserController.Register)
	router.POST("/login", UserController.Login)

	router.POST("/invite", FriendController.Invite)
}

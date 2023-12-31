package router

import (
	"backend/controllers/UserController"
	"backend/database/db"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.GET("/test", func(context *gin.Context) {

		connection := db.PostgresConnection()
		defer connection.Close()

		rows, err := connection.Query("SELECT 121")

		defer rows.Close()

		if err != nil {
			panic(err)
		}

		var tmp int

		for rows.Next() {
			err := rows.Scan(&tmp)

			if err != nil {
				panic(err)
			}

		}

		context.JSON(200, gin.H{
			"message": tmp,
		})
	})

	router.POST("/register", UserController.Register)
	router.POST("/login", UserController.Login)
}

package router

import (
	"backend/database/db"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.GET("/test", func(context *gin.Context) {

		db := db.PostgresConnection()
		defer db.Close()

		rows, err := db.Query("SELECT 121")

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
}

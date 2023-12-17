package cmd

import (
	"backend/router"
	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

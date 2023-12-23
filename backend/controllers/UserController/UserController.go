package UserController

import (
	"backend/controllers/UserController/types"
	"backend/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var form types.RegisterForm
	err := c.BindJSON(&form)
	helpers.HandleError(err)

	fmt.Println(form)
	c.IndentedJSON(200, form)
}

package UserController

import (
	"backend/controllers/UserController/types"
	"backend/database/models/user"
	"backend/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	var form types.RegisterForm
	err := c.Bind(&form)
	helpers.HandleError(err)
	messages, hasError := form.Validate()

	fmt.Println("error: ", err)
	if hasError {
		c.JSON(422, gin.H{
			"message": "error",
			"errors": messages,
		})
		return
	}

	password := helpers.HashPassword(form.Password)

	file, err := c.FormFile("photo")
	helpers.HandleError(err)

	filename := strconv.FormatInt(time.Now().Unix(), 10) + filepath.Ext(file.Filename)

	dst := "storage/user/photos/" + filename
	fmt.Println(dst)
	err = c.SaveUploadedFile(file, dst)
	helpers.HandleError(err)

	newUser := user.User{
		Surname:  form.Surname,
		Name:     form.Name,
		Login:    form.Login,
		Password: password,
		Email:    form.Password,
		Photo:    filename,
	}

	user.CreateUser(&newUser)
	helpers.HandleError(err)

	c.JSON(201, gin.H{
		"message": "created",
	})
}

func Login(c *gin.Context) {
	var form types.LoginForm
	err := c.Bind(&form)
	helpers.HandleError(err)

	currentUser, hasUser := user.FindUserByLogin(form.Login)

	if !hasUser {
		c.JSON(422, gin.H{
			"error": "Пользователя с таким логином не существует",
		})
		return
	}

	if helpers.CheckPassword(form.Password, currentUser.Password) {
		///TODO create JWT
	}

}

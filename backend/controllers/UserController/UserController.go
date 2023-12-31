package UserController

import (
	"backend/controllers/UserController/types"
	"backend/database/models/user"
	"backend/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"path/filepath"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	var form types.RegisterForm
	err := c.Bind(&form)
	if err != nil {
		helpers.ErrorResponse(c, 422, err)
		return
	}

	messages, hasError := form.Validate()

	fmt.Println("error: ", err)
	if hasError {
		c.JSON(422, gin.H{
			"message": "error",
			"errors":  messages,
		})
		return
	}

	password := helpers.HashPassword(form.Password)

	file, err := c.FormFile("photo")
	var filename string

	if file != nil {
		helpers.HandleError(err)

		filename = strconv.FormatInt(time.Now().Unix(), 10) + filepath.Ext(file.Filename)

		dst := "storage/user/photos/" + filename
		fmt.Println(dst)
		err = c.SaveUploadedFile(file, dst)
		helpers.HandleError(err)

	}
	newUser := user.User{
		Surname:  form.Surname,
		Name:     form.Name,
		Login:    form.Login,
		Password: password,
		Email:    form.Email,
		Photo:    filename,
	}

	err = user.CreateUser(&newUser)
	helpers.HandleError(err)

	tokenIdentity := rand.Int()
	payload := map[string]any{
		"token_identity": tokenIdentity,
		"user_id":        newUser.Id,
		"is_admin":       false,
		"exp":            time.Now().Add(time.Minute * 15).Unix(),
	}

	token := helpers.MakeJWT(payload)

	c.SetCookie(
		"access_token",
		token,
		int(time.Now().Add(time.Hour*48).Unix()),
		"/",
		"/",
		true,
		true,
	)

	payload = map[string]any{
		"user_id":        newUser.Id,
		"token_identity": tokenIdentity,
		"exp":            time.Now().Add(time.Hour * 48).Unix(),
	}

	c.JSON(201, gin.H{
		"refresh_token": helpers.MakeJWT(payload),
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
		tokenIdentity := rand.Int()

		payload := map[string]any{
			"user_id":        currentUser.Id,
			"is_admin":       false,
			"token_identity": tokenIdentity,
			"exp":            time.Now().Add(time.Minute * 15).Unix(),
		}

		token := helpers.MakeJWT(payload)

		c.SetCookie(
			"access_token",
			token,
			int(time.Now().Add(time.Hour*48).Unix()),
			"/",
			"/",
			true,
			true,
		)

		payload = map[string]any{
			"user_id":        currentUser.Id,
			"token_identity": tokenIdentity,
			"exp":            time.Now().Add(time.Hour * 48).Unix(),
		}

		c.JSON(201, gin.H{
			"refresh_token": helpers.MakeJWT(payload),
		})
		return
	}

	c.JSON(422, gin.H{
		"error": "Неправильный логин или пароль",
	})
}

func RefreshToken(c *gin.Context)  {
	
}
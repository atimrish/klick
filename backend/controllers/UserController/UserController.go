package UserController

import (
	"backend/controllers/UserController/types"
	"backend/database/models/user"
	"backend/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	JWT "github.com/golang-jwt/jwt/v5"
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

	expAccess := JWT.NumericDate{
		time.Now().Add(time.Minute * 15),
	}

	payload := helpers.CustomClaims{
		UserId:           newUser.Id,
		TokenIdentity:    tokenIdentity,
		IsAdmin:          false,
		RegisteredClaims: JWT.RegisteredClaims{
			ExpiresAt: &expAccess,
		},
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

	expRefresh := JWT.NumericDate{
		time.Now().Add(time.Hour * 48),
	}
	payload = helpers.CustomClaims{
		UserId:           newUser.Id,
		TokenIdentity:    tokenIdentity,
		IsAdmin:          false,
		RegisteredClaims: JWT.RegisteredClaims{
			ExpiresAt: &expRefresh,
		},
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

		expAccess := JWT.NumericDate{
			time.Now().Add(time.Minute * 15),
		}

		payload := helpers.CustomClaims{
			UserId:           currentUser.Id,
			TokenIdentity:    tokenIdentity,
			IsAdmin:          false,
			RegisteredClaims: JWT.RegisteredClaims{
				ExpiresAt: &expAccess,
			},
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

		expRefresh := JWT.NumericDate{
			time.Now().Add(time.Hour * 48),
		}
		payload = helpers.CustomClaims{
			UserId:           currentUser.Id,
			TokenIdentity:    tokenIdentity,
			IsAdmin:          false,
			RegisteredClaims: JWT.RegisteredClaims{
				ExpiresAt: &expRefresh,
			},
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
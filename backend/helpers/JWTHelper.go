package helpers

import (
	"backend/conf"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	JWT "github.com/golang-jwt/jwt/v5"
	"math/rand"
	"time"
)

var jwtSecret = conf.GetConfig()["jwt_secret"]

type CustomClaims struct {
	UserId               int64 `json:"user_id"`
	TokenIdentity        int   `json:"token_identity"`
	IsAdmin              bool  `json:"is_admin"`
	JWT.RegisteredClaims
}

func MakeJWT(payload CustomClaims) string {
	encodedSecret, err := base64.StdEncoding.DecodeString(jwtSecret)
	HandleError(err)
	token := JWT.NewWithClaims(JWT.SigningMethodHS256, payload)

	t, err := token.SignedString(encodedSecret)
	HandleError(err)

	return t
}

func GetPayloadJWT(tokenString string) (*CustomClaims, error) {

	token, err := JWT.ParseWithClaims(tokenString, &CustomClaims{}, func(t *JWT.Token) (interface{}, error) {
		encodedSecret, err := base64.StdEncoding.DecodeString(jwtSecret)
		return encodedSecret, err
	})

	if errors.Is(err, JWT.ErrTokenExpired) {

		return token.Claims.(*CustomClaims), err
	}

	HandleError(err)

	return token.Claims.(*CustomClaims), nil
}

func RefreshToken(accessToken, refreshToken string) (string, string) {
	accessPayload, _ := GetPayloadJWT(accessToken)
	refreshPayload, _ := GetPayloadJWT(refreshToken)

	if accessPayload.TokenIdentity == refreshPayload.TokenIdentity {

		tokenIdentity := rand.Int()

		expAccess := JWT.NumericDate{
			Time: time.Now().Add(time.Minute * 15),
		}

		payload := CustomClaims{
			UserId:        accessPayload.UserId,
			TokenIdentity: tokenIdentity,
			IsAdmin:       false,
			RegisteredClaims: JWT.RegisteredClaims{
				ExpiresAt: &expAccess,
			},
		}

		newAccessToken := MakeJWT(payload)

		expRefresh := JWT.NumericDate{
			Time: time.Now().Add(time.Hour * 48),
		}

		payload.ExpiresAt = &expRefresh

		newRefreshToken := MakeJWT(payload)

		return newAccessToken, newRefreshToken
	}

	return "", ""
}

func TokenExpiredResponse(c *gin.Context) {
	c.JSON(403, gin.H{
		"error": "token expired",
		"message": "TOKEN_EXPIRED_ERROR",
	})
}

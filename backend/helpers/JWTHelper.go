package helpers

import (
	"backend/conf"
	"encoding/base64"
	JWT "github.com/golang-jwt/jwt/v5"
	"math/rand"
	"time"
)

var jwtSecret = conf.GetConfig()["jwt_secret"]

func MakeJWT(payload map[string]any) string {
	encodedSecret, err := base64.StdEncoding.DecodeString(jwtSecret)
	HandleError(err)
	token := JWT.NewWithClaims(JWT.SigningMethodHS256, JWT.MapClaims(payload))

	t, err := token.SignedString(encodedSecret)
	HandleError(err)

	return t
}

func GetPayloadJWT(tokenString string) JWT.MapClaims {
	claims := JWT.MapClaims{}
	_, err := JWT.ParseWithClaims(tokenString, claims, func(t *JWT.Token) (interface{}, error) {
		encodedSecret, err := base64.StdEncoding.DecodeString(jwtSecret)
		return encodedSecret, err
	})
	HandleError(err)

	return claims
}

func RefreshToken(accessToken, refreshToken string) (string, string) {
	accessPayload := GetPayloadJWT(accessToken)
	refreshPayload := GetPayloadJWT(refreshToken)

	if accessPayload["token_identity"] == refreshPayload["token_identity"] {

		tokenIdentity := rand.Int()

		payload := map[string]any{
			"user_id":        accessPayload["user_id"],
			"is_admin":       false,
			"token_identity": tokenIdentity,
			"exp":            time.Now().Add(time.Minute * 15).Unix(),
		}

		newAccessToken := MakeJWT(payload)

		payload = map[string]any{
			"user_id":        accessPayload["user_id"],
			"token_identity": tokenIdentity,
			"exp":            time.Now().Add(time.Hour * 48).Unix(),
		}

		newRefreshToken := MakeJWT(payload)

		return newAccessToken, newRefreshToken
	}

	return "", ""
}

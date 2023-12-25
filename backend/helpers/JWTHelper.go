package helpers

import (
	"backend/conf"
	"github.com/golang-jwt/jwt/v5"
)

var jwt_secret = conf.GetConfig()["jwt_secret"]

func MakeJWT() {

}

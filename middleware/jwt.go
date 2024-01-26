package middleware

import (
	"errors"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func SetToken(username string, password string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		username,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "ginblog",
		},
	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", utils.ERROR
	}
	return token, utils.SUCCESS
}

func CheckToken(token string) (*MyClaims, int) {
	reqClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	}) // 注意，这里的return JwtKey, nil是我们自定义的回调函数的返回值，不是ParseWithClaims的返回值，ParseWithClaims的返回值是*Token, error
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, utils.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
		return nil, utils.ERROR_AUTH_CHECK_TOKEN_FAIL
	}
	if claims, ok := reqClaims.Claims.(*MyClaims); ok && reqClaims.Valid {
		return claims, utils.SUCCESS
	}
	return nil, utils.ERROR_AUTH_CHECK_TOKEN_FAIL
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			utils.ResponseWithMsg(c, utils.ERROR_TOKEN_EXIST)
			c.Abort()
			return
		}
		checkToken := tokenHeader[7:]
		claims, code := CheckToken(checkToken)
		if code == utils.ERROR_AUTH_CHECK_TOKEN_FAIL {
			utils.ResponseWithMsg(c, utils.ERROR_AUTH_CHECK_TOKEN_FAIL)
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt.Time.Unix() {
			utils.ResponseWithMsg(c, utils.ERROR_TOKEN_RUNTIME)
			c.Abort()
			return
		}
		c.Next()
	}
}

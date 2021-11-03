package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-blog/utils/result"
	"net/http"
	"strings"
	"time"
)

type UserClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var signKey = []byte("shanby")

// CreateToken 创建Token值
func CreateToken(username string) (string, result.Code) {
	claims := UserClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "go-blog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		fmt.Println(err)
		return signedToken, result.ERROR
	}
	return signedToken, result.SUCCESS
}

// ParseToken 解析Token
func ParseToken(token string) (*UserClaims, result.Code) {
	parseToken, _ := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	if claims, ok := parseToken.Claims.(*UserClaims); ok && parseToken.Valid {
		return claims, result.SUCCESS
	} else {
		return claims, result.ERROR
	}
}

// JwtToken Jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code = result.SUCCESS
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			code = result.ERROR_TOKEN_EXIST
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": code.GetCodeMessage(),
			})
			ctx.Abort()
			return
		}

		tokenList := strings.SplitN(token, " ", 2)
		if len(tokenList) != 2 && tokenList[0] != "Bearer" {
			code = result.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": code.GetCodeMessage(),
			})
			ctx.Abort()
			return
		}

		parsedToken, ok := ParseToken(tokenList[1])
		if ok == result.ERROR {
			code = result.ERROR_TOKEN_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": code.GetCodeMessage(),
			})
			ctx.Abort()
			return
		}

		if time.Now().Unix() > parsedToken.ExpiresAt {
			code = result.ERROR_TOKEN_RUNTIME
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": code.GetCodeMessage(),
			})
			ctx.Abort()
			return
		}

		if code != result.SUCCESS {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": code.GetCodeMessage(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("username", parsedToken.Username)
		ctx.Next()
	}
}
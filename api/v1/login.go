package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/middleware"
	"go-blog/model"
	"go-blog/utils/result"
	"net/http"
)

func Login(ctx *gin.Context) {
	var user model.User
	var token string
	_ = ctx.ShouldBindJSON(&user)
	code := model.Login(user.Username, user.Password)

	if code == result.SUCCESS {
		token, code = middleware.CreateToken(user.Username)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
		"token": token,
	})
}

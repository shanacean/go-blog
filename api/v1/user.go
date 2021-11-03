package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"
	"go-blog/utils/result"
	"go-blog/utils/validate"
	"net/http"
	"strconv"
)


func CreateUser(ctx *gin.Context) {
	var user model.User
	_ = ctx.ShouldBindJSON(&user)

	msg, vcode := validate.Validate(&user)
	if vcode != result.SUCCESS {
		ctx.JSON(http.StatusOK, gin.H{
			"status": vcode,
			"message": msg,
		})
		return
	}

	code := model.CheckUser(user.Username)
	if code == result.SUCCESS {
		model.AddUser(&user)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
	})
}

func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": result.SUCCESS,
		"message": result.SUCCESS.GetCodeMessage(),
		"data": data,
		"total": total,
	})
}

func EditUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var user model.User
	_ = ctx.ShouldBindJSON(&user)
	code := model.CheckUser(user.Username)
	if code == result.SUCCESS {
		model.EditUser(id, &user)
	}
	if code == result.ERROR_USERNAME_USED {
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
	})
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
	})
}
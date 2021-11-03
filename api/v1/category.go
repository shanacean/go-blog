package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"
	"go-blog/utils/result"
	"net/http"
	"strconv"
)

func CreateCategory(ctx *gin.Context) {
	var c model.Category
	_ = ctx.ShouldBindJSON(&c)
	code := model.CheckUser(c.Name)
	if code == result.SUCCESS {
		model.AddCategory(&c)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
		"data": c,
	})
}

func GetCategory(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetCategory(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": result.SUCCESS,
		"message": result.SUCCESS.GetCodeMessage(),
		"data": data,
		"total": total,
	})
}

func EditCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var c model.Category
	_ = ctx.ShouldBindJSON(&c)
	code := model.CheckUser(c.Name)
	if code == result.SUCCESS {
		model.EditCategory(id, &c)
	}
	if code == result.ERROR_USERNAME_USED {
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
	})
}

func DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteCategory(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
	})
}

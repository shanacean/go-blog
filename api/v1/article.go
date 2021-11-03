package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"
	"net/http"
	"strconv"
)

// CreateArticle 添加文章
func CreateArticle(ctx *gin.Context) {
	var article model.Article
	_ = ctx.ShouldBindJSON(&article)
	code := model.AddArticle(&article)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
		"data": article,
	})
}

// EditArticle 编辑文章
func EditArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var article model.Article
	_ = ctx.ShouldBindJSON(&article)
	code := model.EditArticle(id, &article)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteArticle(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
	})
}


// GetSingleArticle 查询单个文章信息
func GetSingleArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := model.GetSingleArticle(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
		"data": data,
	})
}
// GetArticles 查询所有文章
func GetArticles(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total, code := model.GetArticles(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
		"data": data,
		"total": total,
	})
}

// GetArticlesByCategory 查询分类下所有文章
func GetArticlesByCategory(ctx *gin.Context) {

	cid, _ := strconv.Atoi(ctx.Param("cid"))
	pageSize, _ := strconv.Atoi(ctx.Param("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Param("pagenum"))
	data, total,code := model.GetArticlesByCategory(cid, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
		"data": data,
		"total": total,
	})
}
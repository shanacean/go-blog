package v1

import (
	"github.com/gin-gonic/gin"
	model "go-blog/model"
	"net/http"
)

func Upload(ctx *gin.Context) {
	file, m, _ := ctx.Request.FormFile("file")

	fileSize := m.Size

	url, code := model.UploadFile(file, fileSize)


	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": code.GetCodeMessage(),
		"url": url,
	})
}

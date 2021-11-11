package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func TestFileUpload(t *testing.T) {
	r := gin.Default()
	// 单文件
	r.POST("file", func(context *gin.Context) {
		if file, err := context.FormFile("file"); err != nil {
			context.String(http.StatusInternalServerError, "upload failed")
		} else {
			fmt.Println(file.Filename)
			fmt.Println(file.Header)
			context.SaveUploadedFile(file, fmt.Sprintf("%v_%v", time.Now().Unix(), file.Filename))
			context.String(http.StatusOK, "upload ok")
		}
	})

	// 多文件 默认32M 限制1M
	r.MaxMultipartMemory = 1 << 20
	r.POST("multiFile", func(context *gin.Context) {
		if form, err := context.MultipartForm(); err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		} else {
			files := form.File["files"]
			for _, file := range files {
				if err := context.SaveUploadedFile(file,
					fmt.Sprintf("%v_%v", time.Now().Unix(), file.Filename)); err != nil {
					context.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
					return
				}
			}
			context.String(200, fmt.Sprintf("upload ok %d files", len(files)))
		}

	})
	r.Run()
}

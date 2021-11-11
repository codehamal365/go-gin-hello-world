package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 测试表单参数
func TestFormParam(t *testing.T) {
	r := gin.Default()
	r.POST("/form", func(context *gin.Context) {
		// 字段值
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "admin")
		context.String(http.StatusOK, "username: %s, password: %s", username, password)
	})
	r.Run()
}

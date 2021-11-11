package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestUrlQueryParam(t *testing.T) {

	r := gin.Default()
	// 测试url参数
	testUrlVariables(r)
	r.Run()
}

func testUrlVariables(r *gin.Engine) {
	/**
	不会忽略http://localhost:8080/user?name=
	这里会返回空值
	*/
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
}

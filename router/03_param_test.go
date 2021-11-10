package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testing"
)

func TestParseParam(t *testing.T) {

	r := gin.Default()
	// 测试路径参数
	testPathVariables(r)
	// 测试url参数
	testUrlVariables(r)
	r.Run()
}

func testPathVariables(r *gin.Engine) {
	/**
		可以匹配：
	     /user//
	     /user/xxx/
		不能匹配：
		  /user/xxx
	*/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
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

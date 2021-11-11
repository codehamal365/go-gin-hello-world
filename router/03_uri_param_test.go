package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testing"
)

func TestUriParam(t *testing.T) {

	r := gin.Default()
	// 测试路径参数
	testPathVariables(r)
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

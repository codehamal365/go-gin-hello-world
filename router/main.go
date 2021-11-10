package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	// 监听端口默认为8080
	// 不传 默认是值是 :8080
	r.Run()
}

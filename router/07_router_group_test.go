package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestRouterGroup(t *testing.T) {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/login", handler)
		v1.GET("/form", handler)
	}

	v2 := r.Group("v2")
	{
		v2.GET("/login", handler)
		v2.GET("/form", handler)
	}
	r.Run(":9000")
}

func handler(c *gin.Context) {
	c.String(http.StatusOK, "url:%s", c.Request.URL)
}

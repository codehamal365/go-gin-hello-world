package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestRestRouter(t *testing.T) {
	r := gin.Default()

	r.POST("blogs", add)
	r.DELETE("blogs/:id", deleteById)
	r.PUT("blogs", updateById)
	r.GET("blogs/:id", getById)
	r.GET("blogs", list)

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
	r.Run()
}

func add(context *gin.Context) {
	context.String(http.StatusOK, "blog add")
}

func deleteById(context *gin.Context) {
	context.String(http.StatusOK, "blog delete")
}
func updateById(context *gin.Context) {
	context.String(http.StatusOK, "blog update")
}

func getById(context *gin.Context) {
	context.String(http.StatusOK, "blog get by id")
}
func list(context *gin.Context) {
	context.String(http.StatusOK, "blog list")
}

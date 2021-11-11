package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test404(t *testing.T) {
	r := gin.Default()

	r.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "sorry not found")
	})
	r.Run()
}

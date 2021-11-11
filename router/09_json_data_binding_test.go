package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

// json数据绑定、校验
func TestJsonBinding(t *testing.T) {

	r := gin.Default()
	r.POST("json", func(context *gin.Context) {

		/**
			{
			   "name":"111",
			   "age":14
			}
		  这里 age 不能用字符串包
		*/
		var body loginBody
		if err := context.ShouldBindJSON(&body); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "200", "body": body})
	})
	r.Run()
}

type loginBody struct {
	Name     string    `form:"userName" json:"name" xml:"user" binding:"required" message:"not right"`
	Age      int       `form:"Age" json:"age" xml:"age" binding:"required"`
	BirthDay time.Time `json:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

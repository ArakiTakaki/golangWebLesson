package router

import (
	"github.com/ArakiTakaki/golangWebLesson/controllers/home"

	"github.com/gin-gonic/gin"
)

func homeSet(r *gin.RouterGroup) {
	r.GET("/", home.Index)
	r.GET("/index.html", home.Index)
}

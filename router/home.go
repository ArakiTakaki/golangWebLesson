package router

import (
	"github.com/ArakiTakaki/golangWebLesson/controllers"
	"github.com/gin-gonic/gin"
)

func homeSet(r *gin.RouterGroup) {
	// /home/index.html に飛ぶ様に設定されている。
	r.GET("/index.html", controllers.Index)

}

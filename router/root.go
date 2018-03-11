package router

import (
	"github.com/ArakiTakaki/golangWebLesson/controllers/root"
	"github.com/gin-gonic/gin"
)

func rootSet(r *gin.RouterGroup) {
	// /home/index.html に飛ぶ様に設定されている。（飛ばす場所）
	r.GET("/test.html", root.Index)

}

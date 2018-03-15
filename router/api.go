package router

import (
	"github.com/ArakiTakaki/golangWebLesson/controllers/api"
	"github.com/gin-gonic/gin"
)

func apiSet(r *gin.RouterGroup) {
	// /home/index.html に飛ぶ様に設定されている。（飛ばす場所）
	r.GET("/items", api.NavItems)
	r.GET("/meta", api.PageData)
}

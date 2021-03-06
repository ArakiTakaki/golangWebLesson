package router

import (
	context "github.com/ArakiTakaki/Context"
	"github.com/gin-gonic/gin"
)

// GetRouter ルーターの設定をして、gin.Engineを返す
func GetRouter() *gin.Engine {
	var r = gin.Default()
	context.New()

	// 静的ファイルを配置する場所選択
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/image", "./public/image")
	r.Static("/template", "./public/template")

	//r.LoadHTMLGlob("views/*") 動作しない - 当たり前ではあるが
	r.LoadHTMLGlob("views/**/*")
	rootSet(r.Group(""))
	homeSet(r.Group("/home"))
	apiSet(r.Group("/api"))
	r.NoRoute(func(c *gin.Context) {
		c.HTML(400, "404.html", nil)
	})

	return r
}

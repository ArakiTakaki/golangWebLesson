package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRouter ルーターの設定をして、gin.Engineを返す
func GetRouter() *gin.Engine {
	var r = gin.Default()

	// 静的ファイルを配置する場所選択
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/image", "./public/image")
	r.Static("/jsx", "./public/jsx")

	// ルーティングurlを増やす部分
	var home = r.Group("/home")
	homeSet(home)
	// ルーティングurlを増やす部分

	r.LoadHTMLGlob("views/root/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	return r
}

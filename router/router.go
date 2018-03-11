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

	r.LoadHTMLGlob("views/**/*")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/root/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "root/index.html", nil)
	})
	r.GET("/home/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", nil)
	})
	// ルーティングurlを増やす部分
	//rootSet(r.Group(""))
	//homeSet(r.Group("/home"))
	// ルーティングurlを増やす部分

	return r
}

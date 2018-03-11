package router

import (
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

	//r.LoadHTMLGlob("views/*") 動作しない - 当たり前ではあるが
	r.LoadHTMLGlob("views/**/*")
	//router.LoadHTMLFiles("templates/template1.html","templates/template2.html") // ファイル指定でロード
	rootSet(r.Group(""))
	homeSet(r.Group("/home"))
	r.NoRoute(func(c *gin.Context) {
		c.HTML(400, "404.html", nil)
	})

	return r
}

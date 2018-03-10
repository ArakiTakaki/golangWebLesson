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

	r.LoadHTMLGlob("views/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// // クエリストリングパラメーターは既存のrequest objectの下でパースされる
	// // このリクエストは /welcome?firstname=Jane&lastname=Doe へ返答する
	// r.GET("/welcome", func(c *gin.Context) {
	// 	firstname := c.DefaultQuery("firstname", "Guest") // Geustはデフォルト値?
	// 	lastname := c.Query("lastname")                   // c.Request.URL.Query().Get("lastname") のショートカット

	// 	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	// })
	// // このハンドラは /user/john にマッチするが、 /user/ や /user にはマッチしない
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })

	// // でもこれは /user/john/ や /user/john/send にマッチする
	// //  /user/john にマッチするルーターがなければ、/user/john/ にリダイレクトされる
	// r.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")
	// 	message := name + " is " + action
	// 	c.String(http.StatusOK, message)
	// })

	return r
}

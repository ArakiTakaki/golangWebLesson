package router

import (
	"github.com/ArakiTakaki/golangWebLesson/controllers/home"

	"github.com/gin-gonic/gin"
)

func homeSet(r *gin.RouterGroup) {
	// /home/index.html に飛ぶ様に設定されている。（飛ばす場所）
	r.GET("/", home.Index)
	r.GET("/index.html", home.Index)

}

/*
func main() {
    router := gin.Default()

    // このハンドラは /user/john にマッチするが、 /user/ や /user にはマッチしない
    router.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello %s", name)
    })

    // でもこれは /user/john/ や /user/john/send にマッチする
    //  /user/john にマッチするルーターがなければ、/user/john/ にリダイレクトされる
    router.GET("/user/:name/*action", func(c *gin.Context) {
        name := c.Param("name")
        action := c.Param("action")
        message := name + " is " + action
        c.String(http.StatusOK, message)
    })

    router.Run(":8080")
}
func main() {
    router := gin.Default()

    // クエリストリングパラメーターは既存のrequest objectの下でパースされる
    // このリクエストは /welcome?firstname=Jane&lastname=Doe へ返答する
    router.GET("/welcome", func(c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "Guest") // Geustはデフォルト値?
        lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") のショートカット

        c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
    })
    router.Run(":8080")
}
*/

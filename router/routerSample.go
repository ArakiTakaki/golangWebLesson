package router

import "github.com/gin-gonic/gin"

func TestGetRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("views/**/*")
	/*
		/views/home/index.html
		/views/root/index.html
		ダメだった
		r.GET("/home", func(c *gin.Context) {
			c.HTML(200, "home/index.html", nil)
		})
		r.GET("/root", func(c *gin.Context) {
			c.HTML(200, "root/index.html", nil)
		})
	*/

	r.GET("/home", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	//なぜかコレだったらいける

	return r
}

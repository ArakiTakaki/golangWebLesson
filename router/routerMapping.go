package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Set(r *gin.RouterGroup) {
	r.GET("/test.html", ctl)

}
func ctl(c *gin.Context) {
	render := gin.H{
		"title": "site_title",
	}
	fmt.Println("INDEX表示")
	c.HTML(http.StatusOK, "index.tmpl", render)
}

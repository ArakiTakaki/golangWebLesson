package root

import (
	"github.com/gin-gonic/gin"
)

// Index インデックスの処理系統を記載する。
func Index(c *gin.Context) {
	c.HTML(200, "root/index.tmpl", nil)
}

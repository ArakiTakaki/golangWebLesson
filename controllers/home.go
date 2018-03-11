package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index インデックスの処理系統を記載する。
func Index(c *gin.Context) {
	//ディレクトリを選択するけれど、なるべく１回書くだけで済ませたいところ。
	c.HTML(http.StatusOK, "/home/index.html", nil)
}

package api

import (
	"net/http"

	"github.com/ArakiTakaki/golangWebLesson/conf"
	"github.com/gin-gonic/gin"
)

// Items インデックスの処理系統を記載する。
func Items(c *gin.Context) {
	//var items
	//c.JSON(http.StatusOK, items)
	//ディレクトリを選択するけれど、なるべく１回書くだけで済ませたいところ。
}

// PageData confのメタデータ群を返す、詳細はconfにて
func PageData(c *gin.Context) {
	meta := conf.GetMeta()
	c.JSON(http.StatusOK, meta)
}

func NavItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "test2"})
}

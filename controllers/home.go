package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ctl(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", nil)
}

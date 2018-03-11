package router

import (
	"github.com/gin-gonic/gin"
)

func homeSet(r *gin.RouterGroup) {
	r.GET("/test.html", controllers.Ctr)

}

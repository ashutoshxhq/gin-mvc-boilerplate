package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webkodes.com/admin/controllers"
)

// IndexRouter created auth router
func IndexRouter(r *gin.Engine) *gin.Engine {

	r.GET("/", controllers.Index)

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	return r
}

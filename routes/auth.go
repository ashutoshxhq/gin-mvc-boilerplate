package routes

import (
	"github.com/gin-gonic/gin"
	"webkodes.com/admin/controllers"
)

// AuthRouter created auth router
func AuthRouter(r *gin.Engine) *gin.Engine {

	r.GET("/login", controllers.Login)
	r.GET("/signup", controllers.Signup)

	return r
}

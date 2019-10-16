package main

import (
	"github.com/gin-gonic/gin"
	"webkodes.com/admin/routes"
)

func main() {
	r := gin.Default()
	r.Static("/public", "./public")
	r.LoadHTMLGlob("views/*")
	r = routes.AuthRouter(r)
	r = routes.IndexRouter(r)
	r.Run()
}

package controllers

import "github.com/gin-gonic/gin"

// Index handles login get req
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "index",
	})
}

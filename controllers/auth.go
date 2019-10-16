package controllers

import "github.com/gin-gonic/gin"

// Login handles login get req
func Login(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "login",
		"name":    name,
	})
}

// Signup handles signup get req
func Signup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "signup",
	})
}

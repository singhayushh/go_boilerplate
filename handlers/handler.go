package handlers

import "github.com/gin-gonic/gin"

// RenderHome ... render the index.html page
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Home | Company Name",
	})
}

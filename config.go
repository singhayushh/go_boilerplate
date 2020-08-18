package handlers

import (

)

Router *gin.Engine
func(router *Router) Init(port string) {
	router = gin.New()

	router.initRoutes()

	router.LoadHTMLGlob("views/*.html")
	router.Static("/css", "views/resources/css")
	router.Static("/js", "views/resources/js")
}
package handlers

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
}

func (server *Server) apiStatus() {
	server.Router.GET("/api", Welcome)
}

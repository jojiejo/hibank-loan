package handlers

func (server *Server) GenerateRoutes() {
	api := server.Router.Group("/api")
	api.POST("/loan", server.CreateLoan)
}

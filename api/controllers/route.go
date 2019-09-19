package controllers

import "github.com/tiyodev/api-rest-go-v1/api/middlewares"

func (server *Server) initializeRoutes() {
	// people routes
	server.Router.HandleFunc("/people/{id}", middlewares.SetMiddlewareJSON(server.GetPeople)).Methods("GET")
	server.Router.HandleFunc("/peoples", middlewares.SetMiddlewareJSON(server.GetPeoples)).Methods("GET")
	server.Router.HandleFunc("/people", middlewares.SetMiddlewareJSON(server.CreatePeople)).Methods("POST")
	server.Router.HandleFunc("/people/{id}", middlewares.SetMiddlewareJSON(server.UpdatePeople)).Methods("PUT")
	server.Router.HandleFunc("/people/{id}", middlewares.SetMiddlewareJSON(server.DeletePeople)).Methods("DELETE")
}

package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server data
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// InitializeRoute config
func (server *Server) InitializeRoute() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

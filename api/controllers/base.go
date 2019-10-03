package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/tiyodev/api-rest-go-v2/api/responses"
)

// Server struct
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// InitializeRoute config
func (server *Server) InitializeRoute() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// GetHome : check if the server is working properly
func (server *Server) GetHome(resWriter http.ResponseWriter, req *http.Request) {
	responses.JSON(resWriter, http.StatusOK, "Welcome to this GOLAND REST API")
}

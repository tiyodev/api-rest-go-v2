package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite dialect
	"github.com/tiyodev/api-rest-go-v2/api/controllers"
)

var server = controllers.Server{}

// Run http server
func run(addr string) {
	var errDb error

	// Database connection
	server.DB, errDb = gorm.Open("sqlite3", "db/swapi.dat")
	if errDb != nil {
		log.Fatal(errDb)
	} else {
		fmt.Println("Successful connection to the database")
	}

	// Initialise HTTP router
	server.InitializeRoute()

	fmt.Println("Listening to port 8080")

	// Run HTTP server
	errSrv := http.ListenAndServe(":8080", server.Router)
	if errSrv != nil {
		log.Fatal(errSrv)
	}
}

func main() {
	run(":8080")
}

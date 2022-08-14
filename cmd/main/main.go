package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yega-Noragami/go-chaos/pkg/constants"
	"github.com/Yega-Noragami/go-chaos/pkg/models"
	"github.com/Yega-Noragami/go-chaos/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Setup Connection to Database
	models.SETUP_DB()
	// Setup Router
	r := mux.NewRouter()
	routes.RegisterListRoutes(r)
	http.Handle("/", r)
	// listen and serve will create a server , this is inside the HTTP package on route 9010
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", constants.ServerPort), r))
	fmt.Printf("Server started on %d!", constants.ServerPort)
}

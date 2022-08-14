package routes

import (
	"github.com/Yega-Noragami/go-chaos/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterListRoutes = func(router *mux.Router) {

	router.HandleFunc("/lists", controllers.CreateList).Methods("POST")
	router.HandleFunc("/lists", controllers.GetList).Methods("GET")

	// OPTIONAL
	// router.HandleFunc("/lists/{listkey}", controllers.GetListByKey).Methods("GET")
	// router.HandleFunc("/list/{listkey}", controllers.UpdateList).Methods("PUT")
	// router.HandleFunc("/list/{listkey}", controllers.DeleteList).Methods("DELETE")
}

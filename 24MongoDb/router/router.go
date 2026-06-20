package router

import (
	"github.com/gorilla/mux"
	"github.com/nenri07/mongodb/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/getall", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/create", controller.CreateData).Methods("POST")
	router.HandleFunc("/api/update/{id}", controller.WatchedSwitch).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controller.DeleteMany).Methods("DELETE")

	return router

}

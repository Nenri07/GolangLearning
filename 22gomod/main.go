package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello this is mod class")
	hello()

	r := mux.NewRouter()
	r.HandleFunc("/", show).Methods("GET")

	log.Fatal(http.ListenAndServe(":4340", r))
}

func hello() {
	fmt.Println("Hello")
}
func show(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Nenrei is best</h1>"))

}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nenri07/mongodb/router"
)

func main() {
	fmt.Println("Server is starting.....")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4573", r))
	fmt.Println("Listening on Port 4743....")
}

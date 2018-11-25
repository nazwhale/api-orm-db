package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("servin up some hotcakes...")
	handleRequests()
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "sup wherld")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

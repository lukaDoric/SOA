package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	fmt.Println("Pokrenut staticki web server na port-u 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

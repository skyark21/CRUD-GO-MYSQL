package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Server Listening in Port: 3001")
	addr := ":3001"
	http.ListenAndServe(addr, nil)
	http.HandleFunc("/", Start)
}

func Start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GOLA MUNDO")
}

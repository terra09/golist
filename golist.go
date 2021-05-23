package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "golist server bootstrap!")
	fmt.Println("rootPage has been accessed")
}

func handleRequest() {
	http.HandleFunc("/", rootPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequest()
}

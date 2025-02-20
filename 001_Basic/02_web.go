package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellowFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/", hellowFunc)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

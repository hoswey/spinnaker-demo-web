package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("version")

	if version == "" {
		version = "Unknow"
	}

	fmt.Fprintf(w, "Hi there, Current version is %s!", version)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

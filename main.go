package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	fmt.Println(path)
	_, err := os.Stat(path)
	if err == nil {
		// file exists
		http.ServeFile(w, r, path)
	} else {
		// file does not exists
		index := "index.html"
		_, err := os.Stat(index)
		if err == nil {
			// index.html exists
			http.ServeFile(w, r, index)
		} else {
			// index.html does not exists
			http.NotFound(w, r)
		}
	}
}

func main() {
	port := ":5050"
	http.HandleFunc("/", handler)
	fmt.Println("spa-server starting on localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("spa-server: ", err)
	}
}

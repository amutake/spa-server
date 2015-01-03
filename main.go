package main

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	fmt.Print(r.URL.Path)
	file, err := os.Stat(path)
	if err == nil && !file.IsDir() {
		// file exists
		gz := path + ".gz"
		file, err := os.Stat(gz)
		t := mime.TypeByExtension(filepath.Ext(path))
		if err == nil && !file.IsDir() && t != "" {
			fmt.Println(" => " + gz)
			w.Header().Add("Content-Encoding", "gzip")
			w.Header().Add("Content-Type", t)
			http.ServeFile(w, r, gz)
		} else {
			fmt.Println(" => " + path)
			http.ServeFile(w, r, path)
		}
	} else {
		// file does not exist
		index := "index.html"
		file, err := os.Stat(index)
		if err == nil && !file.IsDir() {
			// index.html exists
			fmt.Println(" => " + index)
			http.ServeFile(w, r, index)
		} else {
			// index.html does not exist
			fmt.Println(" => NotFound")
			http.NotFound(w, r)
		}
	}
}

func main() {

	port := ":5050"
	flag.Parse()
	if flag.NArg() != 0 {
		port = ":" + flag.Arg(0)
	}
	fmt.Println("spa-server starting on localhost" + port)

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("spa-server: ", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		file := filepath.Join("web", filepath.Clean(r.URL.Path[1:]))

		info, err := os.Stat(file)
		if os.IsNotExist(err) || info.IsDir() {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, file)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

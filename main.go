package main

import (
	"errors"
	"log"
	"net/http"
)

const addr = ":3000"

func main() {
	dirName := "/home/aryan/projects/search_engine/dataset/"

	files := ListFiles(dirName)
	for _, v := range files {
		log.Println(v)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", Test)
	mux.HandleFunc("/fetch", CrawlWiki)

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Println("Server started at port: ", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Println("Error running http server: ", err)
		}
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println("Testing server")
}

func CrawlWiki(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://en.wikipedia.org/w/index.php?title=Pet_door&action=raw")
	if err != nil {
		log.Println("Error fetching response from wikipedia")
		log.Fatalln("Error: ", err)
	}

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading body from wiki data")
		log.Fatalln("Error: ", err)
	}

	fmt.Println(string(resData))
}

func (indexes Index) SearchQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if len(query) <= 0 {
		log.Println("No search query provided")
		return
	}
	queries := strings.Split(query, " ")

	for _, query := range queries {
		// start from here later
		// perform intersection of multiple slices
	}

	log.Println("search term: ", query)
	log.Println("Index: ", indexes[query])
}

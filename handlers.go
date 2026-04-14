package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hbollon/go-edlib"
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

func (indexNKeys IndexWithKeys) SearchQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if len(query) <= 0 {
		log.Println("No search query provided")
		return
	}

	queries := strings.Split(query, " ")

	// index1 := indexes["zippy"]
	// index2 := indexes["zvonko"]
	//
	// intersection := Intersection(index1, index2)

	var intersection []int

	startTime := time.Now()

	for _, query := range queries {
		queryIndexes := indexNKeys.indexes[query]
		log.Printf("Number of Indexes found for %v:  %v", query, len(queryIndexes))
		if len(intersection) == 0 {
			intersection = append(intersection, queryIndexes...)
		} else {
			intersection = Intersection(intersection, queryIndexes)
		}
	}

	log.Println("Time taken to search for the word: ", time.Since(startTime))

	log.Println("search term: ", query)
	log.Println("Intersection: ", intersection)

	// log.Println("Index: ", indexes[query])
}

func (indexNKeys IndexWithKeys) SearchSingle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if len(query) <= 0 {
		log.Println("No search query provided")
		return
	}

	// res, err := edlib.FuzzySearchThreshold(query, indexNKeys.keys, 0.5, edlib.Levenshtein)
	res, err := edlib.FuzzySearchSet(query, indexNKeys.keys, 5, edlib.Levenshtein)
	if err != nil {
		log.Panicln("Error implementing fuzzy search: ", err)
	}

	fmt.Println("Result: ", res)
}

func Intersection(a, b []int) []int {
	var i, j int
	result := []int{}
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			result = append(result, a[i])
			i++
			j++
		}
	}
	return result
}

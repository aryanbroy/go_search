package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const addr = ":3000"

func main() {
	dirName := "/home/aryan/projects/search_engine/dataset/"

	// ReadFile("/home/aryan/projects/search_engine/dataset/alt.atheism/54485")
	files := ListFiles(dirName)

	var docIndex int64
	var docs []Document

	for _, v := range files {
		docIndex += 1
		data := ReadFile(v)
		newDoc := Document{
			Id:   docIndex,
			Text: data,
		}
		newDoc.Text = newDoc.RemoveStopWords()

		docs = append(docs, newDoc)
	}

	// docs_enc, err := json.Marshal(docs)
	_, err := json.Marshal(docs)
	if err != nil {
		log.Fatalln("Error marshaling json: ", err)
	}

	// log.Println("Documents: ")
	// log.Println(string(docs_enc))

	invertedIndexes := make(Index, 0)

	for _, doc := range docs {
		tokens := Tokenize(doc.Text)
		for _, token := range tokens {
			doc.InvertedIndexing(token, invertedIndexes)
		}
	}

	// log.Println("Indexes: ", invertedIndexes)

	mux := http.NewServeMux()
	mux.HandleFunc("/", Test)
	mux.HandleFunc("/fetch", CrawlWiki)
	mux.HandleFunc("/search", invertedIndexes.SearchQuery)

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

package main

import (
	// "log"
	"strings"

	"github.com/bbalet/stopwords"
)

func (doc Document) InvertedIndexing(token string, invertedIndexes Index) {
	// log.Println("Token: ", token)
	invertedIndexes[token] = append(invertedIndexes[token], doc.Id)
}

func (doc *Document) RemoveStopWords() string {
	// log.Printf("Removing stopwords from doc %v...", doc.Id)
	cleanStr := stopwords.CleanString(doc.Text, "en", true)
	// log.Println("Stopwords removed")
	return cleanStr
}

func Tokenize(text string) []string {
	lowerCaseText := ToLowerCase(text)
	splitWords := SplitWords(lowerCaseText)
	return splitWords
}

func ToLowerCase(text string) string {
	lowerText := strings.ToLower(text)
	return lowerText
}

func SplitWords(text string) []string {
	s := strings.Fields(text)
	return s
}

package main

import "strings"

func ToLowerCase(text string) string {
	lowerText := strings.ToLower(text)
	return lowerText
}

func SplitWords(text string) []string {
	s := strings.Fields(text)
	return s
}

func Tokenize(text string) []string {
	lowerCaseText := ToLowerCase(text)
	splitWords := SplitWords(lowerCaseText)
	return splitWords
}

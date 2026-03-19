package main

type Document struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
}

type Index map[string][]int64

type InvertedIndexes struct {
	Indexes []Index
}

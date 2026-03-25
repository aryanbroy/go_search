package main

type Document struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type Index map[string][]int

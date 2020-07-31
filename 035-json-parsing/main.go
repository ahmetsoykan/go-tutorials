package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA stroy")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", story["intro"].Paragraphs)
}

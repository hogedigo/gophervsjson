package main

import (
	"encoding/json"
	"log"
	"strings"
)

type Book struct {
	Title   string   `json:"title"`
	Price   int      `json:"price,omitempty"`
	Authors []string `json:"authors"`
}

func main() {
	testJson := `
	{
		"title": "Gopher vs JSON!",
		"price": 1000,
		"authors": [
		"山田　太郎",
			"鈴木　一郎"
		]
	}
	{
		"title": "Go phrasebook",
		"price": 2000,
		"authors": [
			"ほげ　ぼすけて"
		]
	}
	{
		"title": "Go! Gopher",
		"price": 3000,
		"authors": [
			"テスト"
		]
	}
`
	
	dec := json.NewDecoder(strings.NewReader(testJson))
	
	for {
		var b Book
		if err := dec.Decode(&b); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s", b)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Authors []string

func (a *Authors) UnmarshalJSON(b []byte) error {
	s := strings.TrimPrefix(string(b), "\"")
	s = strings.TrimSuffix(s, "\"")
	*a = strings.Split(s, "-")
	return nil
}

type Book struct {
	Title   string  `json:"title"`
	Price   int     `json:"price,omitempty"`
	Authors Authors `json:"authors"`
	Secret  string  `json:"-"`
}

type Bookshelf struct {
	Books []Book `json:"books"`
}

func main() {

	testJson := `
{
	"books": [
		{
			"title": "Gopher vs JSON!",
			"price": 1000,
			"authors": "山田　太郎-鈴木　一郎"
		},
		{
			"title": "Go phrasebook",
			"price": 2000,
			"authors": "ほげ　ぼすけて"
		},
		{
			"title": "Go! Gopher",
			"price": 3000,
			"authors": "テスト"
		},
		{
			"title": "",
			"authors": "hoge"
		}
	]
}
`
	var b Bookshelf
	if err := json.Unmarshal([]byte(testJson), &b); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%s", b)
}

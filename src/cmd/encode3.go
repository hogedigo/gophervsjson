package main

import (
	"encoding/json"
	"log"
)

type Book struct {
	Title   string   `json:"title"`
	Price   int      `json:"price,omitempty"`
	Authors []string `json:"authors"`
	Secret  string   `json:"-"`
}

type Bookshelf struct {
	Books []Book `json:"books"`
}

func main() {
	bs := Bookshelf{
		[]Book{
			{"Gopher vs JSON!", 1000, []string{"山田　太郎", "鈴木　一郎"}, "KKK"},
			{"Go phrasebook", 2000, []string{"ほげ　ぼすけて"}, "xxx"},
			{"Go! Gopher", 3000, []string{"テスト"}, "zzz"},
			{"", 0, []string{"hoge"}, "hoge"},
		},
	}
	b, err := json.MarshalIndent(bs, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", b)
}

package main

import (
	"encoding/json"
	"log"
)

type Book struct {
	Title   string
	Price   int
	Authors []string
}

type Bookshelf struct {
	Books   []Book
}

func main(){
	bs := Bookshelf{
		[]Book{
			{"Gopher vs JSON!", 1000, []string{"山田　太郎", "鈴木　一郎"}},
			{"Go phrasebook", 2000, []string{"ほげ　ぼすけて"}},
			{"Go! Gopher", 3000, []string{"テスト"}},
		},
	}
	b, err := json.MarshalIndent(bs, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", b)
}


package main

import (
	"encoding/json"
	"log"
	"os"
)

type Book struct {
	Title   string   `json:"title"`
	Price   int      `json:"price,omitempty"`
	Authors []string `json:"authors"`
}

func main() {
	enc := json.NewEncoder(os.Stdout)
	
	if err := enc.Encode(&Book{"Gopher vs JSON!", 1000, []string{"山田　太郎", "鈴木　一郎"}}); err != nil {
		log.Fatal(err)
	}
	if err := enc.Encode(&Book{"Go phrasebook", 2000, []string{"ほげ　ぼすけて"}}); err != nil {
		log.Fatal(err)
	}
	if err := enc.Encode(&Book{"Test", 2000, []string{"ZZZZ"}}); err != nil {
		log.Fatal(err)
	}
}

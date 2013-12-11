package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	Title   string
	Price   int
	Authors []string
}

func main() {

	book := Book{"Gopher vs JSON!", 2000, []string{"yamada taro", "suzuki ichiro"}}
	b, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}


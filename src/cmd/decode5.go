package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	testJson := `
{
	"books": [
		{
			"title": "Gopher vs JSON!",
			"price": 1000,
			"authors": [
				"山田　太郎",
				"鈴木　一郎"
			]
		},
		{
			"title": "Go phrasebook",
			"price": 2000,
			"authors": [
				"ほげ　ぼすけて"
			]
		},
		{
			"title": "Go! Gopher",
			"price": 3000,
			"authors": [
				"テスト"
			]
		},
		{
			"title": "",
			"authors": [
				"hoge"
			]
		}
	]
}
`
	var b interface{}
	if err := json.Unmarshal([]byte(testJson), &b); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%s\n", b)
	
	bookshelf := b.(map[string]interface{})
	books := bookshelf["books"].([]interface{})
	book2 := books[2].(map[string]interface{})

	fmt.Printf("%s\n", book2["title"])
}

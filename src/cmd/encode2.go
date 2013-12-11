package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	bs := map[string]interface{}{
		"Books": []map[string]interface{}{
			map[string]interface{}{
				"Title":   "Gopher vs JSON!",
				"Price":   1000,
				"Authors": []string{"山田　太郎", "鈴木　一郎"},
			},
			map[string]interface{}{
				"Title":   "Go phrasebook",
				"Price":   2000,
				"Authors": []string{"山田　太郎", "鈴木　一郎"},
			},
			map[string]interface{}{
				"Title":   "Go! Gopher",
				"Price":   3000,
				"Authors": []string{"テスト"},
			},
		},
	}
	b, err := json.MarshalIndent(bs, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}

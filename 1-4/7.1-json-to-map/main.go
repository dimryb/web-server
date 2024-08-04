package main

import (
	"encoding/json"
	"fmt"
)

var JSON_STRING = `
{
	"id": 55,
	"price": 3000,
	"items": [
		{
			"name": "snowboard",
			"number": 1
		},
		{
			"name": "ball",
			"number": 4
		}
	]
}
`

type Data struct {
	Price int `json:"price"`
}

func main() {
	var data map[string]any

	err := json.Unmarshal([]byte(JSON_STRING), &data)
	if err != nil {
		panic(err)
	}

	x, _ := data["id"].(string)


	fmt.Printf("%+v", x)
}
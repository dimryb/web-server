package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int	`json:"id"`
	Name string `json:"username"`
}

func main() {
	user1 := User{Name: "Ivan", Id: 555}
	bytes, _ := json.Marshal(user1)
	fmt.Println(string(bytes))
	fmt.Println(bytes)
	fmt.Println(user1)
}
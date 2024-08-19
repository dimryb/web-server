package main

import (
	"fmt"
	"time"
)

func main() {
	printString("string0")
	go printString("string1")
	go printString("string2")
	printString("string3")
	time.Sleep(1 * time.Second)
}

func printString(str string) {
	for i := range(10){
		fmt.Println(str, i)
		time.Sleep(1 * time.Microsecond)
	}	
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	s := make(chan string)
	var wg sync.WaitGroup
	var wgOut sync.WaitGroup

	wgOut.Add(1)
	go func() {
		defer wgOut.Done()
		fmt.Println("output")
		for sOut := range s {
			fmt.Println(sOut)
		}
	}()

	printString("string0", s)

	wg.Add(1)
	go func() {
		defer wg.Done()
		printString("string1", s)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		printString("string2", s)
	}()

	printString("string3", s)

	wg.Wait()
	close(s)

	wgOut.Wait()
}

func printString(str string, s chan string) {
	fmt.Println("<-", str)
	s <- fmt.Sprint("chain ", str)
}

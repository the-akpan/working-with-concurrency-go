package main

import (
	"fmt"
	"sync"
)

var msg string
var msgs = []string{
	"Hello, universe!",
	"Hello, cosmos!",
	"Hello, world!",
}

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	msg = "Hello, world!"

	for _, word := range msgs {
		wg.Add(1)
		go updateMessage(word, &wg)
		wg.Wait()
		printMessage()
	}
}

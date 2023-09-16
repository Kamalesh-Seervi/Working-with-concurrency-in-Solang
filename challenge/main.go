package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	msg = "Hello, world!"
	wg.Add(3)

	go func() {
		updateMessage("Hello, universe!")
		printMessage()
		defer wg.Done()
	}()

	go func() {
		updateMessage("Hello, world!")
		printMessage()
		defer wg.Done()
	}()

	go func() {
		updateMessage("Hello, cosmos!")
		printMessage()
		defer wg.Done()
	}()
	wg.Wait()
}

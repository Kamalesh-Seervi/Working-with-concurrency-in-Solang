package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var msg string

func UpdateMessage(s string) {
	defer wg.Done()
	msg = s
}

func PrintMessage() {
	fmt.Println(msg)
}

func main() {
	msg = "Kamalesh"
	wg.Add(2)
	go UpdateMessage("Hello")
	go UpdateMessage("Kamalesh")
	wg.Wait()
	PrintMessage()

}

// func UpdateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()
// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

// func PrintMessage() {
// 	fmt.Println(msg)
// }

// func main() {
// 	msg = "Kamalesh"
// 	var mutex sync.Mutex
// 	wg.Add(2)
// 	go UpdateMessage("Hello", &mutex)
// 	go UpdateMessage("Kamalesh", &mutex)
// 	wg.Wait()
// 	PrintMessage()

// }

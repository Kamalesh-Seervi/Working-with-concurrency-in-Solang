package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Add(a int, b int) int {
	return a + b
}

func main() {

	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := Add(4, 5)
			fmt.Println(result)
		}()
	}

	wg.Wait()

}

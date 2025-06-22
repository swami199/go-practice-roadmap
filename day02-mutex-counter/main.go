package main

import (
	"fmt"
	"sync"
)

var Counter int
var Mutex sync.Mutex

func counterIncrement(wg *sync.WaitGroup) {
	// defer fmt.Println("Done with counter")
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		Mutex.Lock()
		Counter++
		Mutex.Unlock()
	}
	// fmt.Println(Counter)
}

func main() {
	// var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go counterIncrement(&wg)
	}
	wg.Wait()
	fmt.Println("final output ", Counter)

}

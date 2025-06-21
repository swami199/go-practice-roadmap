package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d work has started\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d work has finished\n", id)

}

func main() {
	var wg sync.WaitGroup
	numberOfWorkers := 5
	wg.Add(numberOfWorkers)

	for i := 0; i <= 5; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
}

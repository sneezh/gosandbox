package main

import (
	"sync"
)

// исходная задача
func main_() {
	for i := 0; i < 10; i++ {
		go func() {
			println(i)
		}()
	}
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println(i)
		}(i)
		wg.Wait()
	}
}

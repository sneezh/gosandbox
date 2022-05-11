package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println(1)
	time.Sleep(1 * time.Second)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println(2)
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		fmt.Println(3)
		wg.Done()
	}()
	wg.Wait()
}

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go walk(1)
	go walk(2)
	go walk(3)
	go walk(4)

	time.Sleep(50 * time.Microsecond)
}

func walk(leg int) {
	for {
		fmt.Println(leg)
		runtime.Gosched()
	}
}

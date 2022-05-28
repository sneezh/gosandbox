package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	m := make(map[string]string, 100)
	rwm := &sync.RWMutex{}
	for i := 1; i <= 100; i++ {
		writeToMap(rwm, m, fmt.Sprintf("%v", i), fmt.Sprintf("%v", i+rand.Intn(500)))
	}
	for i := 1; i <= 100; i++ {
		go func(i int) {
			r := readFromMap(rwm, m, fmt.Sprintf("%v", i))
			fmt.Println(r)
		}(i)
	}
	for i := 1; i <= 100; i++ {
		go writeToMap(rwm, m, fmt.Sprintf("%v", i), fmt.Sprintf("%v", i+rand.Intn(500)))
	}
	for i := 1; i <= 100; i++ {
		go writeToMap(rwm, m, fmt.Sprintf("%v", i), fmt.Sprintf("%v", i+rand.Intn(500)))
	}
	for i := 1; i <= 100; i++ {
		go func(i int) {
			r := readFromMap(rwm, m, fmt.Sprintf("%v", i))
			fmt.Println(r)
		}(i)
	}
	for i := 1; i <= 100; i++ {
		go func(i int) {
			r := readFromMap(rwm, m, fmt.Sprintf("%v", i))
			fmt.Println(r)
		}(i)
	}

	time.Sleep(2 * time.Second)
}

func readFromMap(rwm *sync.RWMutex, m map[string]string, key string) string {
	rwm.RLock()
	v := m[key]
	rwm.RUnlock()
	return v
}

func writeToMap(rwm *sync.RWMutex, m map[string]string, key, value string) {
	rwm.Lock()
	m[key] = value
	rwm.Unlock()
}

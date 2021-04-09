package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var m  = sync.RWMutex{}

var storage int32 = 0

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go printStorage()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func printStorage() {
	fmt.Printf("Storage value: %v\n", storage)
	m.RUnlock()
	wg.Done()
}

func increment() {
	storage++
	m.Unlock()
	wg.Done()
}
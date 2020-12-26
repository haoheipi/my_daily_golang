package main

import (
	"fmt"
	"sync"
	"time"
)

var m1 sync.Mutex
var m2 sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		m1.Lock()
		time.Sleep(1000)
		m2.Lock()
		fmt.Println("do something")
		m2.Unlock()
		m1.Unlock()
	}()

	go func() {
		defer wg.Done()
		m2.Lock()
		time.Sleep(1000)
		m1.Lock()
		fmt.Println("do something")
		m1.Unlock()
		m2.Unlock()
	}()

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	balance := 100
	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {

		mutex.Lock()
		balance = balance + 1000
		fmt.Println(balance, "Depositer")
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		mutex.Lock()
		balance = balance - 1000
		fmt.Println(balance, "Withdrawer")
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		mutex.Lock()
		fmt.Println(balance, "Balance Checker 1")
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		mutex.Lock()
		fmt.Println(balance, "Balancer Checker 2")
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
}

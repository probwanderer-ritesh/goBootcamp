package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func deposit(amount int64, balance *int64, mu *sync.Mutex) {

	mu.Lock()
	defer mu.Unlock()
	atomic.AddInt64(balance, amount)

}

func withdraw(amount int64, balance *int64, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	atomic.AddInt64(balance, -amount)
}

func getBalance(balance *int64, mu *sync.Mutex) int64 {
	return atomic.LoadInt64(balance)
}

func main() {
	var wg sync.WaitGroup
	var balance int64
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			deposit(100, &balance, &mu)
		}()
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			withdraw(20, &balance, &mu)
		}()
	}

	wg.Wait()

	fmt.Println("Final Balance:", getBalance(&balance, &mu))
}

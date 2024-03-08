package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func getRating(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var rating int = rand.Intn(6-1) + 1
	fmt.Println(rating)
	ch <- rating

}
func main() {
	var totalRatingSum float64 = 0
	var wg sync.WaitGroup
	ch := make(chan int, 200)
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go getRating(ch, &wg)
	}
	wg.Wait()
	close(ch)
	for i := 0; i < 200; i++ {
		totalRatingSum += float64(<-ch)
	}
	fmt.Println(totalRatingSum / 200)
}

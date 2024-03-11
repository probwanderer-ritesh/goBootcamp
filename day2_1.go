package goBootcamp

import (
	"fmt"
	"strings"
	"sync"
)

func count(s string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	chars := strings.Split(s, "")
	var count = make(map[string]int)
	//fmt.Println(s)
	for i := 0; i < len(chars); i++ {
		count[chars[i]] = count[chars[i]] + 1
	}

	fmt.Println("frequencies of letters in word ", s, " are:", count)
	ch <- "done"
}

func Driver() {
	var wg sync.WaitGroup
	var s = []string{"quick", "brown", "fox", "lazy", "dog"}
	ch := make(chan string, len(s))
	//[“quick”, “brown”, “fox”, “lazy”, “dog”]
	for i := 0; i < len(s); i++ {
		wg.Add(1)
		go count(s[i], &wg, ch)
	}
	wg.Wait()

}

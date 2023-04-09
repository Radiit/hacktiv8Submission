package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []string{"bisa1", "bisa2", "bisa3"}
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(4)
	c := make(chan int, 4)
	for i := 1; i <= 4; i++ {
		c <- i
		go printData2(&wg, c, &mu)
		fmt.Println(a, i)
	}
	close(c)
	wg.Wait()
}

func printData2(wg *sync.WaitGroup, c chan int, mu *sync.Mutex) {
	b := []string{"coba1", "coba2", "coba3"}
	for x := range c {
		mu.Lock()
		fmt.Println(b, x)
		wg.Done()
		mu.Unlock()
	}
}

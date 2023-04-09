package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []string{"bisa1", "bisa2", "bisa3"}
	var wg sync.WaitGroup
	wg.Add(4)
	c := make(chan int, 4)
	for i := 1; i <= 4; i++ {
		c <- i
		go printData2(&wg, c)
		fmt.Println(a, i)
	}
	close(c)
	wg.Wait()
}

func printData2(wg *sync.WaitGroup, c chan int) {
	b := []string{"coba1", "coba2", "coba3"}
	for x := range c {
		fmt.Println(b, x)
		wg.Done()
	}
}

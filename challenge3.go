package main
package main

import "fmt"

func main() {
	var a string = "selamat malam"
	print(a)
	maps := count(a)
	fmt.Printf("map[%c:%d", ' ', maps[' '])
	for i := 'a'; i <= 'z'; i++ {
		if maps[i] > 0 {
			fmt.Printf(" %c:%d", i, maps[i])
		}
	}
	fmt.Printf("]\n")
}

func print(a string) {
	for _, v := range a {
		fmt.Printf("%c\n", v)
	}
}

func count(a string) map[rune]int {
	Maps := make(map[rune]int)
	for _, char := range a {
		Maps[char]++
	}
	return Maps
}

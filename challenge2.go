package main

import "fmt"

func main() {
	var huruf [7]rune = [7]rune{'C', 'A', 'M', 'A', 'P', 'B', '0'}
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}
	for j := 0; j < 11; j++ {
		fmt.Println("Nilai j = ", j)
		if j == 4 {
			for h := 0; h < len(huruf); h++ {
				fmt.Printf("Character %U '%c' starts at byte position %d\n", huruf[h], huruf[h], h*2)
			}
			j = 5
		}
	}
}

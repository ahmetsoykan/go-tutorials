package main

import (
	"fmt"
)

func main() {
	fmt.Println("Verilen sayilarin en buyugunu bulup dondurecektir.")
	greatest := max(2, 6, 3, 76, 123, 62, 12, 79, 124, 1, 502)
	fmt.Println("En buyuk sayi sudur", greatest)
}

func max(numbers ...int) int {
	var v int
	for _, item := range numbers {
		if item > v {
			v = item
			fmt.Println("Gecici en buyuk sayi sudur:", v)
		}
	}
	return v
}

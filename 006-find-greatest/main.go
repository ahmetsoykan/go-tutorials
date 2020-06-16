package main

import (
	"fmt"
)

func main() {
	fmt.Println("Verilen sayilarin en buyugunu bulup dondurecektir.")
	greatest := max(2, 6, 4, 76, 126, 62, 12, 76, 124, 13, 501)
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

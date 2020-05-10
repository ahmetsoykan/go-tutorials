package main

import (
	"fmt"
)

func main() {
	sntnc := reverse("ahmet soykan!")
	for _, let := range sntnc {
		defer fmt.Printf(let)
	}
}

func reverse(s string) []string {
	sntnc := []rune(s)
	list := make([]string, len(sntnc))
	for i, let := range sntnc {
		list[i] = string(let)
	}
	return list
}

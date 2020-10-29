package main

import "fmt"

func main() {

	givenArray := []int{2, 1, 5, 7, 4}
	stdArray := make([]int, 0)

	for x, _ := range givenArray {
		stdArray = append(stdArray, x)
	}
	for i, _ := range givenArray {
		if stdArray[i] == givenArray[i] {
			fmt.Printf("%v indexli item: %v\n", i, givenArray[i])
			break
		}
	}
}

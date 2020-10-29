package main

import "fmt"

func main() {

	givenArray := []int{2, 1, 5, 7, 4}

	for i, item := range givenArray {
		if item == i {
			fmt.Printf("%v indexli item: %v\n", i, item)
			break
		}
	}

}

package main

import "fmt"

func main() {

	d := []int{13, 12, 48, 1, 4, -4, 2, -3, 20, 101, 0, 100, 91, 99, 11, 23, 3, 7, -6}

	sorted := Sort(d)

	fmt.Println("Sorted slice:", sorted)

}

func Sort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}

		}
	}

	return arr
}

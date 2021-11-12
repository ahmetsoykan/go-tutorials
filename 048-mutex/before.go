// // Go is a language known for how simple it is to run concurrent routunes.
// // However it's easy to run into errors when concurrent goroutines have to access the same piece of data.
// // Mutexes are data structures provided by the sync package. They can help us prevent concurrent access to data that we want to change atomically.
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func isEven(n int) bool {
// 	return n%2 == 0
// }

// func main() {
// 	n := 0

// 	// goroutune 1
// 	go func() {
// 		nIsEven := isEven(n)
// 		time.Sleep(5 * time.Millisecond)
// 		if nIsEven {
// 			fmt.Println(n, " is even")
// 			return
// 		}
// 		fmt.Println(n, "is odd") // reading n from here after goroutine2 has changed the value of n
// 	}()
// 	// goroutine 2
// 	go func() {
// 		n++
// 	}()

// 	// just waiting for the gouritines to finish before exiting
// 	time.Sleep(time.Second)

// }

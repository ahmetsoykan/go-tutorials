package main

import "fmt"

func square(ch chan int) {
	for item := 1; item <= 10; item++ {
		ch <- item * item
	}
	close(ch)
}

func main() {
	c := make(chan int)

	go square(c)

	for val := range c {
		fmt.Println(val)
	}
}

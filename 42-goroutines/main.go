package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("Series 1")
		time.Sleep(2 * time.Second)
		done <- true
	}()
	fmt.Println("Series 2")
	<-done
}

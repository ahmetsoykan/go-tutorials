package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	for {
		file, _ := ioutil.ReadFile("./config")
		keys := strings.Split(string(file), "=")

		for i := 0; i < len(keys); i++ {
			fmt.Println(keys[i])
		}

		time.Sleep(1 * time.Second)
	}
}

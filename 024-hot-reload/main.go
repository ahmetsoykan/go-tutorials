package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Config map[string]string

func ReadConfig(filename string) Config {

	config := Config{}

	file, _ := os.Open(filename)

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
	}
	return config
}

func main() {
	for {
		conf := ReadConfig(`./config`)
		fmt.Println(conf["NAME"])

		time.Sleep(1 * time.Minute)
	}
}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Dosya acilamadi!")
	}

	r := csv.NewReader(file)
	all, _ := r.ReadAll()

	for i := range all {
		key := strings.Split(all[i][0], "=")
		value := strings.Split(all[i][0], "=")

		fmt.Printf("%d.nci satir { key: %s, value: %s }\n", i+1, key[0], value[1])
	}
}

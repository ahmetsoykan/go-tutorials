package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("keys.csv")
	if err != nil {
		fmt.Println("Dosya acilamadi!")
	}

	r := csv.NewReader(file)
	all, _ := r.ReadAll()

	source := make(map[string]string)
	for i := range all {
		key := strings.Split(all[i][0], "=")
		value := strings.Split(all[i][0], "=")
		source[key[0]] = value[1]

		fmt.Printf("%d.nci satir { key: %s, value: %s }\n", i+1, key[0], value[1])
	}

	fmt.Println(source["NAME"])
	fmt.Println(source["SURNAME"])
	fmt.Println(source["NUMBER"])
}

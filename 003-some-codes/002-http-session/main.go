package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("./sales.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}

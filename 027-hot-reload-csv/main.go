package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("config.csv")
	r := csv.NewReader(file)
	lines, _ := r.ReadAll()

	line := parseLines(lines)
	fmt.Print(line[0].v)
}

func parseLines(lines [][]string) []keys {
	ret := make([]keys, len(lines))
	for i, line := range lines {
		ret[i] = keys{
			k: line[0],
			v: line[1],
		}
	}
	return ret
}

type keys struct {
	k string
	v string
}

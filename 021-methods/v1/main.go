package main

import (
	"fmt"
)

// Struct type - `Point`
type Point struct {
	X, Y float64
}

// Method with receiver `Point`
func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}

func main() {
	p := Point{2.0, 3.0}

	fmt.Println("Point : ", p)
	fmt.Println("Is point p located aboce the line y = 1.0 ? : ", p.IsAbove(1))
}

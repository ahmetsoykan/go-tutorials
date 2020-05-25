package main

import (
	"fmt"
	"math"
)

type AritmeticProgression struct {
	A, D float64
}

// Method with receiver `AritmeticProgression`
func (ap AritmeticProgression) NthTerm(n int) float64 {
	return ap.A + float64(n-1)*ap.D
}

type GeometricProgression struct {
	A, R float64
}

// Method with receiver `GeometricProgression`
func (gp GeometricProgression) NthTerm(n int) float64 {
	return gp.A * math.Pow(gp.R, float64(n-1))
}

func main() {
	ap := AritmeticProgression{1, 2}
	gp := GeometricProgression{1, 2}

	fmt.Println("5th Term of the Arithmetic series = ", ap.NthTerm(5))
	fmt.Println("5th Term of the Geometric series = ", gp.NthTerm(5))
}

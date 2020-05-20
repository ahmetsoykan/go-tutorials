package main

// Sum calculates the total from a slice of numbers
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll calculates the respective sums of every slice passed in
func SumAll(numbersToSum ...[]int) []int {
	lenghtOfNumberts := len(numbersToSum)
	sums := make([]int, lenghtOfNumberts)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

package main

import "fmt"

var table = []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func words(digits string) (result []string) {
	helper(&result, digits, "", 0)
	return
}
func helper(result *[]string, digits, prefix string, index int) {
	if index == len(digits) {
		*result = append(*result, prefix)
		return
	}
	digit := digits[index] - '0'
	for _, v := range table[digit] {
		helper(result, digits, prefix+string(v), index+1)
	}

}
func main() {
	fmt.Println(words("534"))
}
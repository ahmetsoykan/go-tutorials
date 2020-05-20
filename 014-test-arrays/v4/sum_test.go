package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collections of any size", func(t *testing.T) {
		numberts := []int{1, 2, 3}

		got := Sum(numberts)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numberts)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})

	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

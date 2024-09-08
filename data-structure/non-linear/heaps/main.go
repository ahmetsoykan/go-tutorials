package main

import "fmt"

// heaps
// remember heap sort
// easy to implement priority queues

// Max key is the root
// gimme the highest key
// it looks like binary tree

// time complexity depend on height of the tree
// O(h)
// O(log n) becuase it's binary tree

// Heapify
// Extract

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array is empty")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when the left is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}

	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 7, 9, 11, 13, 15, 17}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	fmt.Println("insert done")

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
	fmt.Println("extract done")

}

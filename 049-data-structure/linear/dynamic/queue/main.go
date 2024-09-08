package main

import "fmt"

// Stack and Queue
// Difference is how you take data out

// Queues are FIFO (First in first out)
// Queues use ENQUEUE and DEQUEUE

// Queue represents stack that hold a slice
type Queue struct {
	items []int
}

// Enqueue add a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue remove first value at the queue
// and Returns the removed value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.Enqueue(100)
	fmt.Println(myQueue)

	myQueue.Enqueue(200)
	fmt.Println(myQueue)

	myQueue.Enqueue(300)
	fmt.Println(myQueue)

	myQueue.Dequeue()
	fmt.Println(myQueue)

}

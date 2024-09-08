package main

import "fmt"

// Linked List
// Like Array but every item linked to next one
// It has HEAD, NEXT...
// Easy to add and removing values at the begining of the list
// When we want to add item to array, all the values are shifhted to right.

// Doubly Linked List
// Each Node has NEXT and PREVIOUS

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

// as it's just mutating data, it can take pointer receiver
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// as it's just printing data, it can take value receiver
func (l linkedList) printListData() {
	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {

	// empty list
	if l.length == 0 {
		return
	}

	// if the value matches with head
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {

	mylinkedList := linkedList{}

	node1 := &node{
		data: 48,
	}
	node2 := &node{
		data: 18,
	}
	node3 := &node{
		data: 16,
	}

	node4 := &node{
		data: 11,
	}

	node5 := &node{
		data: 7,
	}
	node6 := &node{
		data: 2,
	}

	mylinkedList.prepend(node1)
	mylinkedList.prepend(node2)
	mylinkedList.prepend(node3)
	mylinkedList.prepend(node4)
	mylinkedList.prepend(node5)
	mylinkedList.prepend(node6)

	fmt.Println(mylinkedList)

	mylinkedList.printListData()

	mylinkedList.deleteWithValue(16)
	mylinkedList.printListData()

}

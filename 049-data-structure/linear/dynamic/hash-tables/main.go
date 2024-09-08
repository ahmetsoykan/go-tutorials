package main

import "fmt"

// key=value
// we can this this as a combination of array + linked list
// array but index utilize hash functions
// hash functions
// if we have 7 sized array, we put input hash function and get base 7, then result the index
// collision handling, 2 input might select the same index
// we might put second input the next index
// but we lose benefits if we do this
// seperate chainging play a role then
// then we can put these item to linked list and we put this linkedlist pointer in the index
// time complexity is constant O(1)

// Maps are hashtables in Go

// ArraySize is the size of the hash table array
const arraySize int = 7

// Hashtable structure
// HashTable will hold an array
type Hashtable struct {
	array [arraySize]*bucket
}

// Bucket structure
// bucket is a linked list in each slof ot the
type bucket struct {
	head *bucketNode
}

// bucketNode structure
// bucketNode is a linked list node that hold the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *Hashtable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

// Search will take in a key and return true if that key is stored in the hash table
func (h *Hashtable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *Hashtable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert will take in key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("item exists")
	}
}

// search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false

}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {
	previousNode := b.head

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *Hashtable {
	result := &Hashtable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func main() {

	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KELLY",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	fmt.Println(hashTable)

	hashTable.Delete("STAN")

	fmt.Println(hashTable)

	fmt.Println(hashTable.Search("KELLY"))

}

package main

import "fmt"

// Binary Tree
// Root, Left Children, Right Children
// Left is smaller, right is bigger
// When the tree is balanced, advantage is speed.
// O(log n)

// Node represent the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {

	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			// recursive call
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left

		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			// recursive call
			n.Left.Insert(k)
		}
	}

}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {

	// end of the tree
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}
	return true
}

func main() {

	// tree is just a struct literal
	tree := &Node{Key: 100}

	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(19)
	tree.Insert(150)
	tree.Insert(76)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(276)

	// should return 300
	fmt.Println(tree.Right.Right)

	// should return true
	fmt.Println(tree.Search(276))
	// should return false
	fmt.Println(tree.Search(2176))

}

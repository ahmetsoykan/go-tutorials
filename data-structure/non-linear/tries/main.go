package main

import "fmt"

// TRIES
// Works better with words search like as you type
// It has Root, Child
// isEnd = False|True
// time complexity is O(m) - m is the lenght of the word
// but sometimes it's faster if it's not matched at the begining when going through tris

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}

		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search will take in a word and RETURN true is that word is included in the trie
func (t *Trie) Search(w string) bool {

	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false

}

func main() {
	myTrie := InitTrie()

	toAdd := []string{
		"aragorn",
		"aragon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("aragorn"))
}

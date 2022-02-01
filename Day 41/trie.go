/*
O(m) time complexity
*/
package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node
}

//Create a new trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//Insert
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		// w[i] - 'a' because that way 'a' becomes 0, 'b' 1 and so on.
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		// w[i] - 'a' because that way 'a' becomes 0, 'b' 1 and so on.
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
	testTrie := InitTrie()
	testTrie.Insert("vasilis")
	fmt.Println(testTrie.Search("vasilis"))

}

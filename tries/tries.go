package main

import "fmt"

//
const AlphabetSize = 26

// Node represents each node in the tries
type Node struct {
	childern [26]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.childern[charIndex] == nil {
			currentNode.childern[charIndex] = &Node{}
		}
		currentNode = currentNode.childern[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and RETURN true is that word is included in the
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.childern[charIndex] == nil {
			return false
		}
		currentNode = currentNode.childern[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

func main() {
	testTrie := InitTrie()
	toAdd := []string{
		"aragon",
		"nilesh",
		"neel",
		"riddhesh",
	}

	for _, v := range toAdd {
		testTrie.Insert(v)
	}
	fmt.Println(testTrie.Search("ne"))
}

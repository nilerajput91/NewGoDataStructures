package main

import "fmt"

const ArraySize = 7

// HashTable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is linked list each slot of the array
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node  that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Search will take in a key and return true if that key is stored  in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

// Delete will take in a key and delete it from the hashtable
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	preivousNode := b.head
	for preivousNode != nil {
		if preivousNode.next.key == k {
			// delete
			preivousNode.next = preivousNode.next.next
		}
		preivousNode = preivousNode.next

	}
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// Insert will take in linked list node that holds the key
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// insert will take in a key ,create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// search will take in a key and return true if the bucket  has that key
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

// hash function  to calculte the hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize

}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	fmt.Println(hashTable)
	fmt.Println(hashTable.Search("ERIC"))
	hashTable.Delete("RANDY")
}

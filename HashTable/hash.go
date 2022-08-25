package hashtable

import (
	"fmt"
)

// collision handling - open addressing and separate chaining(used)
const arraySize = 7

// HashTable structure
type Hashtable struct {
	array [arraySize]*bucket
}

// Bucket structure
type bucket struct {
	head *bucketNode
}

// bucket node structure
type bucketNode struct {
	next *bucketNode
	key  string
}

// Insert --> Insert the ket into hashtable
func (h *Hashtable) Insert(key string) {
	if h.Search(key) {
		fmt.Println("key already exists")
		return
	} else {
		index := hash(key)
		h.array[index].insert(key)
	}
}

// Search --> checks if key exits int table
func (h *Hashtable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete
func (h *Hashtable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert
func (b *bucket) insert(k string) {
	node := &bucketNode{key: k}
	node.next = b.head
	b.head = node
}

// search
func (b *bucket) search(key string) bool {
	temp := b.head
	for temp != nil {
		if temp.key == key {
			return true
		}
		temp = temp.next
	}
	return false
}

// delete
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	node := b.head
	for node.next != nil {
		if node.next.key == k {
			node.next = node.next.next
			return
		}
		node = node.next
	}
}

// hash function
func hash(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % 7
}

// init
func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

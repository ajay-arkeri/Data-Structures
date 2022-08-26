package main

import "fmt"

type node struct {
	key   int
	left  *node
	right *node
}

// insert
func (l *node) Insert(k int) {
	if l.key < k { //right
		if l.right == nil {
			l.right = &node{key: k}
		} else {
			l.right.Insert(k)
		}
	} else { //right
		if l.left == nil {
			l.left = &node{key: k}
		} else {
			l.left.Insert(k)
		}
	}
}

// search
func (l *node) Search(k int) bool {
	if l == nil {
		return false
	}
	if l.key < k { //right
		return l.right.Search(k)
	} else if l.key > k { //left
		return l.left.Search(k)
	}
	return true //match
}

//delete
func (l *node) Delete(key int) {

}

func main() {
	t := &node{key: 100}
	s := []int{1, 2, 3, 4, 5, 6, 7, 90}
	for _, num := range s {
		t.Insert(num)
	}
	fmt.Println(t.Search(1))
}

//Leaf node --> no children

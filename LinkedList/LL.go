package main

import (
	"fmt"
)

type LinkedList struct {
	head   *node
	length int
}

type node struct {
	next *node
	data any //any alias for interface{}
}

func (l *LinkedList) Prepend(val any) {
	new := &node{data: val}
	temp := l.head
	l.head = new
	l.head.next = temp
	l.length++
}

func (l *LinkedList) DeleteFromVal(val any) {
	if l.length == 0 {
		fmt.Println("empty list")
		return
	}

	if l.head.data == val {
		l.head = l.head.next
	}
	temp := l.head

	for temp.next != nil {
		if temp.next.data == val {
			temp.next = temp.next.next
			l.length--
			return
		}
		temp = temp.next
	}
	fmt.Println("key not found")
}

func (l LinkedList) Println() {
	for l.head != nil {
		fmt.Println(l.head.data)
		l.head = l.head.next
	}
}

// inbuilt ---> conatiner/list

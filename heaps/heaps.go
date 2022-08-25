package heaps

import "fmt"

//Use - priority queues
//a complete tree - all levels are full except the lowest level (where some nodes can be empty but all nodes to the left)
//But stored as an array
//Max Heap -- Max stored in root and every parent higher than its children

//MaxHeap struct which holds a slice
type MaxHeap struct {
	heapArray []int
}

//Insert Method
func (h *MaxHeap) Insert(key int) {
	h.heapArray = append(h.heapArray, key)
	h.maxheapifyUp(len(h.heapArray) - 1)
}

//Extract maximum method
func (h *MaxHeap) Extract() int {
	extractedValue := h.heapArray[0]
	if len(h.heapArray) == 0 {
		fmt.Println("cannot extract from empty heap")
		return -1
	}
	h.heapArray[0] = h.heapArray[len(h.heapArray)-1]
	h.heapArray = h.heapArray[:len(h.heapArray)-1]

	h.maxheapifyDown(0)

	return extractedValue
}

func (h *MaxHeap) maxheapifyDown(index int) {
	childToCompare := 0
	l, r := left(index), right(index)
	lastIndex := len(h.heapArray) - 1
	for l <= lastIndex {
		//if only left child
		if l == lastIndex {
			childToCompare = l
		} else if h.heapArray[l] > h.heapArray[r] {
			childToCompare = l // if left is greater
		} else {
			childToCompare = r // if right is greater
		}

		if h.heapArray[index] < h.heapArray[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) maxheapifyUp(index int) {
	for h.heapArray[parent(index)] < h.heapArray[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.heapArray[i1], h.heapArray[i2] = h.heapArray[i2], h.heapArray[i1]
}

package main

import (
	"container/heap"
	"fmt"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	hLen := len(*h)
	x := (*h)[hLen-1]
	*h = (*h)[:hLen-1]
	return x
}

func (h Heap) Top() int {
	return h[0]
}

type MinHeap struct {
	Heap
}

func (h MinHeap) Less(i, j int) bool {
	return h.Heap[i] < h.Heap[j]
}

type MaxHeap struct {
	Heap
}

func (h MaxHeap) Less(i, j int) bool {
	return h.Heap[i] > h.Heap[j]
}

func main() {
	minH := &MinHeap{}
	heap.Init(minH)
	heap.Push(minH, 5)
	fmt.Println(minH.Top())
	// minH := &MinHeap{Heap{5, 7, 3, 1}}
	// heap.Init(minH)
	// heap.Push(minH, 4)
	// fmt.Printf("MinHeap Top: %d\n", minH.Top())             // Output: MinHeap Top: 1
	// fmt.Printf("Popped from MinHeap: %d\n", heap.Pop(minH)) // Output: Popped from MinHeap: 1
	// fmt.Printf("New MinHeap Top: %d\n", minH.Top())         // Output: New MinHeap Top: 3
	// fmt.Println("Hi V ", minH.Heap[0])
	// maxH := &MaxHeap{Heap{1, 3, 5, 7}}
	// heap.Init(maxH)
	// heap.Push(maxH, 6)
	// fmt.Printf("MaxHeap Top: %d\n", maxH.Top())             // Output: MaxHeap Top: 7
	// fmt.Printf("Popped from MaxHeap: %d\n", heap.Pop(maxH)) // Output: Popped from MaxHeap: 7
	// fmt.Printf("New MaxHeap Top: %d\n", maxH.Top())         // Output: New MaxHeap Top: 6
}

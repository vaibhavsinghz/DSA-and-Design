package main

import (
	"container/heap"
	"fmt"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool { // comparator
	return h[i] < h[j]
	/*
		Building minHeap here
		We check if in min heap the h[i] is less than h[j], that means we are in correct order
		as h[i] should be smaller than h[j] for minHeap
	*/
}

func (h Heap) Swap(i, j int) {
	//swapping to implement when comparator gives false
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	hLen := len(*h)
	x := (*h)[hLen-1]  // last element
	*h = (*h)[:hLen-1] // removed last index element
	return x
}

func main() {
	list := Heap{7, 1, 9, 4, 5}
	heap.Init(&list)
	heap.Push(&list, 9)
	fmt.Println(heap.Pop(&list))
	fmt.Println(list)
	fmt.Println(heap.Pop(&list))
	fmt.Println(list)
	fmt.Println(heap.Pop(&list))
	fmt.Println(list)
}

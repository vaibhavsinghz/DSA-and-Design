package main

import (
	"container/heap"
	"fmt"
	"math"
)

type pair struct {
	first, second int
}

type minHeap []pair

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].second < h[j].second
}

func (h minHeap) Swap(i, j int) {
	h[i].first, h[i].second = h[j].first, h[j].second
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *minHeap) Pop() any {
	hLen := len(*h)
	x := (*h)[hLen-1]
	*h = (*h)[:hLen-1]
	return x
}

func getPair(v, w int) pair {
	return pair{
		first:  v,
		second: w,
	}
}
func dijkstra(V, src int, adj [][][]int) []int {
	costTo := make([]int, V)
	for i := range costTo {
		costTo[i] = math.MaxInt
	}
	costTo[src] = 0

	pq := &minHeap{}
	heap.Init(pq)
	pq.Push(getPair(src, 0))

	for pq.Len() > 0 {
		cur := pq.Pop().(pair)
		curNode, curWt := cur.first, cur.second

		for _, nbr := range adj[curNode] {
			v, w := nbr[0], nbr[1]

			if curWt+w < costTo[v] {
				costTo[v] = curWt + w
				pq.Push(getPair(v, costTo[v]))
			}
		}
	}

	return costTo
}

func main() {
	V, S := 3, 2
	adj := make([][][]int, V)
	adj[0] = append(adj[0], []int{1, 1})
	adj[0] = append(adj[0], []int{2, 6})
	adj[1] = append(adj[1], []int{2, 3})
	adj[1] = append(adj[1], []int{0, 1})
	adj[2] = append(adj[2], []int{1, 3})
	adj[2] = append(adj[2], []int{0, 6})

	dijkstra(V, S, adj)

	help()
	// fmt.Println(res)
}

func isOutside(i, j, r, c int) bool {
	if i < 0 || i >= r || j < 0 || j >= c {
		return true
	}
	return false
}

func getVertexNo(i, j, c int) int {
	return i*c + j
}

func help() {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	grid := [][]int{
		// {0, 1, 0, 0, 0},
		// {0, 1, 0, 1, 0},
		// {0, 0, 0, 1, 0},
		{0, 1, 1},
		{1, 1, 0},
		{1, 1, 0},
	}
	r, c := len(grid), len(grid[0])
	V := r * c
	adj := make([][][]int, V)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			curVertexNo := getVertexNo(i, j, c)
			for _, dir := range dirs {
				newI, newJ := i+dir[0], j+dir[1]
				if isOutside(newI, newJ, r, c) {
					continue
				}
				nbrVertexNo := getVertexNo(newI, newJ, c)
				adj[curVertexNo] = append(adj[curVertexNo], []int{nbrVertexNo, grid[newI][newJ]})
			}
		}
	}
	fmt.Println("---------")
	for i := range adj {
		fmt.Println(adj[i])
	}
	fmt.Println("---------")
	res := dijkstra(V, 0, adj)
	fmt.Println(res)
}

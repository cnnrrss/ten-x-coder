package main

import (
	"container/heap"
	"fmt"
)

func main() {
	ans := findKthLargest([]int{3, 4, 2, 1, 5, 66, 7, 9, 144, 2324, 99}, 4)
	fmt.Println(ans)
}

func findKthLargest(nums []int, k int) int {
	temp := PriorityQueue(nums)
	pq := &temp
	heap.Init(pq)

	if k == 1 {
		return (*pq)[0]
	}

	for i := 1; i < k; i++ {
		fmt.Println(heap.Remove(pq, 0))
	}
	return (*pq)[0]
}

// An Item is something we manage in a priority queue.
// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []int

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] > pq[j] }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return item
}

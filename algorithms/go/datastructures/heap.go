package main

import "container/heap"

type minHeap []int

func findKthLargest(nums []int, k int) int {
	temp := minHeap(nums)
	h := &temp
	heap.Init(h)
	if k == 1 {
		return (*h)[0]
	}
	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}

	return (*h)[0]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}

func (h *minHeap) Len() int           { return len(h) }
func (h *minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

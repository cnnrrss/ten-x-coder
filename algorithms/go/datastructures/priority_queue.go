package main

import "container/heap"

type pair struct {
    i, j int
    sum int
}

type priorityQueue []*pair

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].sum < pq[j].sum }
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j], pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) {
	p := x.(*pair)
	*pq = append(*pq, p)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	pair := old[n-1]
	*pq = old[0:n-1]
	return p
}

func kSmallestPairs(pairs1, pairs2 []int, k int) [][]int {
	var result [][]int
	if len(pairs1) == 0 || len(pairs2) == 0 {
		return res
	}

	pqLen := minimum(k, len(pairs1))
	for l := 0; l < k && l < len(a); l++ {
		pq[l] = &pair{i: l, j: 0, sum: a[l]+b[0]}
	}
	heap.Init(&pq)

	var min *pair
	for ; k > 0 && len(pq) > 0; k-- {
		min = heap.Pop(&pq).(*pair)
		res = append(res, []int{pair1[min.i], pair2[min.j]})
		if min.j+1 < len(pair2) {
			heap.Push(&pq, &pair{i: min.i, j:min.j+1, sum: pair1[min.i] + pair2[min.j+1]})
		}
	}
}

func minimum(a, b int) int {
	if a < b{
		return a
	}
	return b
}
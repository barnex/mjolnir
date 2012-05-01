package helheim

import (
//"container/heap"
)

type PriorityQueue []*Job

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	//pq[i].index = i
	//pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	// To simplify indexing expressions in these methods, we save a copy of the
	// slice object. We could instead write (*pq)[i].
	a := *pq
	n := len(a)
	a = a[0 : n+1]
	item := x.(*Job)
	//item.index = n
	a[n] = item
	*pq = a
}

func (pq *PriorityQueue) Pop() interface{} {
	a := *pq
	n := len(a)
	item := a[n-1]
	//item.index = -1 // for safety
	*pq = a[0 : n-1]
	return item
}

// update is not used by the example but shows how to take the top item from
// the queue, update its priority and value, and put it back.
//func (pq *PriorityQueue) update(value string, priority int) {
//    item := heap.Pop(pq).(*Job)
//    item.value = value
//    item.priority = priority
//    heap.Push(pq, item)
//}

// changePriority is not used by the example but shows how to change the
// priority of an arbitrary item.
//func (pq *PriorityQueue) changePriority(item *Job, priority int) {
//    heap.Remove(pq, item.index)
//    item.priority = priority
//    heap.Push(pq, item)
//}

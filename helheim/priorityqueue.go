package helheim

import ()

// used internally by JobQueue.
type priorityQueue []*Job

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].priority > pq[j].priority {
		return true
	}
	return pq[i].id < pq[j].id
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	// To simplify indexing expressions in these methods, we save a copy of the
	// slice object. We could instead write (*pq)[i].
	a := *pq
	item := x.(*Job)
	a = append(a, item)
	*pq = a
}

func (pq *priorityQueue) Pop() interface{} {
	a := *pq
	n := len(a)
	item := a[n-1]
	//item.index = -1 // for safety
	*pq = a[0 : n-1]
	return item
}

// changePriority is not used by the example but shows how to change the
// priority of an arbitrary item.
//func (pq *priorityQueue) changePriority(item *Job, priority int) {
//    heap.Remove(pq, item.index)
//    item.priority = priority
//    heap.Push(pq, item)
//}

package main

import (
	"sync"
)

func (pq *priorityQueue) extractMin() *vertex {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	t := pq.data[0]
	pq.swap(0, pq.size-1)
	pq.size--
	pq.minHeapify(0, pq.size)
	return t
}

func (pq *priorityQueue) update(v *vertex) {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	pq.fix(v.index)
}

func (pq *priorityQueue) fix(i int) {
	pq.minHeapify(i, pq.size)
	p := parent(i)
	for i > 0 && pq.less(i, p) {
		pq.swap(i, p)
		i = p
		p = parent(i)
	}
}

func (pq *priorityQueue) insert(t *vertex) {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	t.index = pq.size
	pq.data = append(pq.data, t)
	pq.fix(pq.size)
	pq.size++
}

func (pq *priorityQueue) isEmpty() bool {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	return pq.size == 0
}

func (pq *priorityQueue) init() {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	n := len(pq.data)
	pq.size = n
	for i := n / 2; i >= 0; i-- {
		pq.minHeapify(i, n)
	}
}

func (pq *priorityQueue) minHeapify(i, n int) {
	l := left(i)
	r := right(i)
	largest := i
	if l < n && pq.less(l, largest) {
		largest = l
	}
	if r < n && pq.less(r, largest) {
		largest = r
	}
	if largest != i {
		pq.swap(largest, i)
		pq.minHeapify(largest, n)
	}
}

func (pq *priorityQueue) swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
	pq.data[i].index = i
	pq.data[j].index = j
}

func (pq *priorityQueue) less(i, j int) bool {
	return pq.data[i].less(pq.data[j])
}

type priorityQueue struct {
	lock sync.Mutex
	data []*vertex
	size int
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

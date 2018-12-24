package main

import (
	"sync"
)

func main() {

}

type queue struct {
	lock sync.Mutex
	q    []item
	size int
}

func (q *queue) enqueue(t item) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.q = append(q.q, t)
	q.size++
}

func (q *queue) dequeue() (item, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		return item{}, false
	}
	result := q.q[0]
	q.q = q.q[1:]
	q.size--
	return result, true
}

func (q *queue) peek() (item, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		return item{}, false
	}
	return q.q[0], true
}

func (q *queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.size == 0
}

type item struct {
	Val int
}

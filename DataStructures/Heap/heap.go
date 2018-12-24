package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int31n(300)
	var sign int64
	A := make([]*item, n)
	var h heap
	for i := int32(0); i < n; i++ {
		if rand.Float32() >= 0.5 {
			sign = -1
		} else {
			sign = 1
		}
		t := &item{priority: sign * rand.Int63n(150)}
		A[i] = t
		h.insert(t)
	}
	heapSort(A)
	var strBuilder strings.Builder
	strBuilder.WriteString("[")
	for _, v := range h.data {
		strBuilder.WriteString(fmt.Sprintf("%d, ", v.priority))
	}
	fmt.Println(strBuilder.String()[:strBuilder.Len()-2] + "]")
}

func heapSort(A []*item) {
	h := heap{data: A}
	h.init()
	h.sort()
}

func (h *heap) update(i int, t *item) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.data[i] = t
	h.fix(i)
}

func (h *heap) fix(i int) {
	h.maxHeapify(i, h.size)
	p := parent(i)
	for i > 0 && h.more(i, p) {
		h.swap(i, p)
		i = p
		p = parent(i)
	}
}

func (h *heap) insert(t *item) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.data = append(h.data, t)
	h.fix(h.size)
	h.size++
}

func (h *heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *heap) sort() {
	h.lock.Lock()
	defer h.lock.Unlock()

	n := h.size - 1
	for i := n; i > 0; i-- {
		h.swap(0, i)
		n--
		h.maxHeapify(0, n)
	}
}

func (h *heap) init() {
	h.lock.Lock()
	defer h.lock.Unlock()

	n := len(h.data)
	h.size = n
	for i := n / 2; i >= 0; i-- {
		h.maxHeapify(i, n)
	}
}

func (h *heap) maxHeapify(i, n int) {
	l := left(i)
	r := right(i)
	largest := i
	if l < n && h.more(l, largest) {
		largest = l
	}
	if r < n && h.more(r, largest) {
		largest = r
	}
	if largest != i {
		h.data[largest], h.data[i] = h.data[i], h.data[largest]
		h.maxHeapify(largest, n)
	}
}

func (h heap) more(i, j int) bool {
	return h.data[i].priority >= h.data[j].priority
}

type heap struct {
	lock sync.Mutex
	data []*item
	size int
}

type item struct {
	val      int
	priority int64
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

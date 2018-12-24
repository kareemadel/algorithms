package main

import "sync"

func main() {

}

type stack struct {
	lock sync.Mutex
	data []item
	size int
}

func (s *stack) push(t item) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, t)
	s.size++
}

func (s *stack) pop() (item, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		return item{}, false
	}
	t := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return t, true
}

func (s *stack) peek() (item, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		return item{}, false
	}
	return s.data[s.size-1], true
}

func (s *stack) isEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.size == 0
}

type item struct {
	Val int
}

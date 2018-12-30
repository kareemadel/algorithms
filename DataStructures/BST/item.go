package bst

func (t *item) greater(s *item) bool {
	return t.Val > s.Val
}

func (t *item) equal(s *item) bool {
	return t.Val == s.Val
}

func (t *item) swap(s *item) {
	t.Val, s.Val = s.Val, t.Val
}

func (t *item) destroy() {
	t.Val = 0
}

type item struct {
	Val int
}

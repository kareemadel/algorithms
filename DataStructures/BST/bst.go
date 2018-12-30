package bst

import "fmt"

func (r *node) getRank(t *item) uint {
	if r == nil || t == nil {
		return 0
	} else if t.equal(r.data) {
		return r.leftSize + 1
	}
	var rank uint
	if t.greater(r.data) {
		rank = r.leftSize + r.count + r.right.getRank(t)
	} else {
		rank = r.left.getRank(t)
	}
	return rank
}

func (r *node) selectOrder(order uint) *node {
	if r == nil || order > r.leftSize && order <= r.leftSize+r.count {
		return r
	}
	if order > r.leftSize {
		return r.right.selectOrder(order - r.leftSize - r.count)
	}
	return r.left.selectOrder(order)
}

func (r *node) delete(count uint) {
	if r == nil {
		return
	} else if count < r.count && count > 0 {
		r.count -= count
		r.reduceParentsSize(count, nil)
		return
	}
	if r.deleteChildless() || r.deleteOneChildNode() {
		return
	}
	r.deleteTwoChildNode()
}

func (r *node) deleteTwoChildNode() bool {
	if r.left == nil || r.right == nil {
		return false
	}
	r.reduceParentsSize(r.count, nil)
	predecessor := r.findPre()
	predecessor.reduceParentsSize(predecessor.count, r.parent)
	r.swapValue(predecessor)
	if predecessor.right == nil {
		predecessor.parent.left = nil
	} else {
		predecessor.parent.left = predecessor.right
		predecessor.right.parent = predecessor.parent
	}
	return true
}

func (r *node) deleteOneChildNode() bool {
	hasRight := r.right != nil
	hasLeft := r.left != nil
	var child *node
	if hasLeft && hasRight || !hasLeft && !hasRight {
		return false
	} else if hasLeft {
		child = r.left
	} else {
		child = r.right
	}
	r.reduceParentsSize(r.count, nil)
	position := r.whichChild()
	if position == 1 {
		r.parent.right = child
	} else if position == -1 {
		r.parent.left = child
	}
	child.parent = r.parent
	return true
}

func (r *node) deleteChildless() bool {
	if r.left != nil || r.right != nil {
		return false
	}
	r.reduceParentsSize(r.count, nil)
	position := r.whichChild()
	if position == 1 {
		r.parent.right = nil
	} else if position == -1 {
		r.parent.left = nil
	}
	r.destroy()
	return true
}

func (r *node) whichChild() int {
	if r.parent == nil {
		return 0
	}
	if r.greater(r.parent) {
		return 1
	}
	return -1
}

func (r *node) reduceParentsSize(count uint, limit *node) {
	movingPtr := r
	var position int
	for movingPtr != limit {
		position = movingPtr.whichChild()
		// fmt.Println(movingPtr.data, position)
		if position == 1 {
			movingPtr.parent.rightSize -= count
		} else if position == -1 {
			movingPtr.parent.leftSize -= count
		} else {
			break
		}
		movingPtr = movingPtr.parent
	}
}

func Print(r *node) {
	r.traverseInOrder()
	fmt.Printf("\n")
}

func (r *node) traverseInOrder() {
	if r == nil {
		return
	}
	r.left.traverseInOrder()
	r.print()
	fmt.Printf("(%v) ", r.count)
	r.right.traverseInOrder()
}

func (r *node) findPre() *node {
	// if it has left branch, then return the maximum of the left subtree
	if r.left != nil {
		return r.left.max()
	}
	// else return the first parent that's less than the node
	movingPtr := r.parent
	for movingPtr != nil && movingPtr.greater(r) {
		movingPtr = movingPtr.parent
	}
	return movingPtr
}

func (r *node) max() *node {
	if r == nil {
		return nil
	}
	movingPtr := r
	for movingPtr.right != nil {
		movingPtr = movingPtr.right
	}
	return movingPtr
}

func (r *node) min() *node {
	if r != nil {
		for r.left != nil {
			r = r.left
		}
	}
	return r
}

func (r *node) search(t *item) *node {
	if r == nil || t.equal(r.data) {
		return r
	}
	if t.greater(r.data) {
		return r.right.search(t)
	}
	return r.left.search(t)
}

func (r *node) insert(t *item) *node {
	k := &node{data: t, count: 1}
	if r == nil {
		return k
	}
	movingPtr := r
	var parent *node
	var isRight bool
	for movingPtr != nil && !movingPtr.equal(k) {
		parent = movingPtr
		if k.greater(movingPtr) {
			movingPtr.rightSize++
			movingPtr = movingPtr.right
			isRight = true
		} else {
			movingPtr.leftSize++
			movingPtr = movingPtr.left
			isRight = false
		}
	}
	if movingPtr == nil {
		movingPtr = k
		movingPtr.parent = parent
		if isRight {
			parent.right = movingPtr
		} else {
			parent.left = movingPtr
		}
	} else {
		movingPtr.count++
	}
	movingPtr.leftSize, movingPtr.rightSize = 0, 0
	return movingPtr
}

func (r *node) destroy() {
	r.data.destroy()
	r.count, r.leftSize, r.rightSize = 0, 0, 0
	r.left, r.right, r.parent = nil, nil, nil
}

func (r *node) swap(k *node) {
	r.swapValue(k)
	r.swapSizes(k)
	r.swapPointers(k)
}

func (r *node) swapPointers(k *node) {
	r.left, k.left = k.left, r.left
	r.right, k.right = k.right, r.right
	r.parent, k.parent = k.parent, r.parent
}

func (r *node) swapSizes(k *node) {
	r.leftSize, k.leftSize = k.leftSize, r.leftSize
	r.rightSize, k.rightSize = k.rightSize, r.rightSize
}

func (r *node) swapValue(k *node) {
	r.data.swap(k.data)
	r.count, k.count = k.count, r.count
}

func (r *node) print() {
	fmt.Printf("%v", r.data.Val)
}

func (r *node) greater(m *node) bool {
	return r.data.greater(m.data)
}

func (r *node) equal(m *node) bool {
	return r.data.equal(m.data)
}

type node struct {
	left      *node
	right     *node
	parent    *node
	data      *item
	count     uint
	leftSize  uint
	rightSize uint
}

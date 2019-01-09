package main

func (n *node) find() *node {
	if n.parent != n {
		n.parent = n.parent.find()
	}
	return n.parent
}

func (n *node) union(m *node) *node {
	nRoot := n.find()
	mRoot := m.find()
	root := nRoot
	if nRoot.rank < mRoot.rank {
		nRoot.parent = mRoot
		root = mRoot
	} else if mRoot.rank < nRoot.rank {
		mRoot.parent = nRoot
	} else {
		mRoot.parent = nRoot
		nRoot.rank++
	}
	return root
}

type node struct {
	label  string
	rank   int
	parent *node
}

type edge struct {
	u      string
	v      string
	weight int
}

package main

func (v *vertex) addEdge(label string, weight uint) {
	if v.edges == nil {
		v.edges = make(map[string]uint)
	}
	v.edges[label] = weight
}

func (v *vertex) addReverseEdge(label string, weight uint) {
	if v.reversedEdges == nil {
		v.reversedEdges = make(map[string]uint)
	}
	v.reversedEdges[label] = weight
}

func (v *vertex) importEdges(u *vertex) {
	for key, val := range u.edges {
		v.edges[key] = val
	}
	for key, val := range u.reversedEdges {
		v.edges[key] = val
	}
}

type vertex struct {
	label         string
	edges         map[string]uint
	reversedEdges map[string]uint
	distance      uint
	prevLabel     string
	index         int
}

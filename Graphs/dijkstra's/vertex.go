package main

func (v *vertex) addEdge(label string, weight uint) {
	if v.edges == nil {
		v.edges = make(map[string]uint)
	}
	v.edges[label] = weight
}

func (v *vertex) importEdges(u *vertex) {
	for key, val := range u.edges {
		v.edges[key] = val
	}
}

type vertex struct {
	label     string
	edges     map[string]uint
	distance  uint
	prevLabel string
	index     int
}

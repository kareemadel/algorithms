package main

func (v *vertex) addEdge(label string, weight int) {
	if v.edges == nil {
		v.edges = make(map[string]int)
	}
	elem, ok := v.edges[label]
	if ok {
		if weight < elem {
			v.edges[label] = weight
		}
	} else {
		v.edges[label] = weight
	}
}

func (v *vertex) importEdges(u *vertex) {
	for key, val := range u.edges {
		v.edges[key] = val
	}
}

func (v *vertex) less(k *vertex) bool {
	return v.prevWeight <= k.prevWeight
}

type vertex struct {
	label      string
	edges      map[string]int
	prevLabel  string
	prevWeight int
	index      int
}

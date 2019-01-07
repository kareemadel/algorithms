package main

import "fmt"

func main() {
	g := graph{}
	g.init()
	g.readFromFile("/home/kareemadel/algorithms/Graphs/MST/prim/edges.txt")
	_, sum := g.getMST()
	fmt.Println(sum)
}

func (g *graph) getMST() (*graph, int) {
	var sum int
	var sLabel string
	var s *vertex
	for k, v := range g.adList {
		sLabel, s = k, v
		break
	}
	mst := graph{}
	mst.init()
	pq := priorityQueue{}
	infinity := int(^uint(0) >> 1)
	s.prevWeight = 0
	i := 0
	for _, v := range g.adList {
		if v.label != sLabel {
			v.prevWeight = infinity
		}
		v.index = i
		pq.data = append(pq.data, v)
		i++
	}
	pq.init()
	for !pq.isEmpty() {
		u := pq.extractMin()
		mst.createVertex(u.label)
		mst.createVertex(u.prevLabel)
		sum += u.prevWeight
		mst.addEdge(u.label, u.prevLabel, u.prevWeight)
		u.index = -1
		for vLabel := range u.edges {
			v := g.getVertex(vLabel)
			if v.index != -1 {
				w := u.edges[vLabel]
				if w < v.prevWeight {
					v.prevWeight = w
					v.prevLabel = u.label
					pq.update(v)
				}
			}
		}
	}
	return &mst, sum
}

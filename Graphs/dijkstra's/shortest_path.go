package main

import "fmt"

func main() {
	g := &graph{}
	g.init()
	g.readFromFile("/home/kareemadel/algorithms/Graphs/dijkstra's/dijkstraData.txt")
	g.dijkstra("1")
	t := []string{"7", "37", "59", "82", "99", "115", "133", "165", "188", "197"}
	fmt.Println("Shortest distance from vertex 1.")
	fmt.Println("--------------------------------")
	for _, v := range t {
		fmt.Printf("Vertex: %s, Distance: %d\n", v, g.getVertex(v).distance)
	}
}

func (g *graph) dijkstra(s string) {
	pq := priorityQueue{}
	infinity := ^uint(0)
	sVertex := g.getVertex(s)
	sVertex.distance = 0
	i := 0
	for _, v := range g.adList {
		if v.label != s {
			v.distance = infinity
		}
		v.index = i
		pq.data = append(pq.data, v)
		i++
	}
	pq.init()
	for !pq.isEmpty() {
		u := pq.extractMin()
		u.index = -1
		for vLabel := range u.edges {
			v := g.getVertex(vLabel)
			if v.index != -1 {
				d := u.distance + u.edges[vLabel]
				if d < v.distance {
					v.distance = d
					v.prevLabel = u.label
					pq.update(v)
				}
			}
		}
	}
}

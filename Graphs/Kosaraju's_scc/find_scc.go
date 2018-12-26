package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	uuid "github.com/satori/go.uuid"
)

func main() {
	g := &graph{}
	g.init()
	gReverse := g.readFromFile("/home/kareemadel/algorithms/Graphs/Kosaraju's_scc/SCC.txt")
	var L []string
	explored := make(map[string]bool)
	label, _ := uuid.NewV4()
	componentName := label.String()
	for _, n := range gReverse.adList {
		L = visit(gReverse, n, explored, L)
	}
	assigned := make(map[string]bool)
	components := make(map[string][]string)
	for i := len(L) - 1; i >= 0; i-- {
		n := g.getNode(L[i])
		if !assigned[n.label] {
			assign(g, n, assigned, components, componentName)
			label, _ := uuid.NewV4()
			componentName = label.String()
		}
	}
	minInt := -(int((^uint(0)) >> 1)) - 1
	maxSCC := make([]int, 5)
	for i := 0; i < len(maxSCC); i++ {
		maxSCC[i] = minInt
	}
	for _, v := range components {
		l := len(v)
		for i := 0; i < len(maxSCC); i++ {
			if l > maxSCC[i] {
				for j := len(maxSCC) - 1; j > i; j-- {
					maxSCC[j] = maxSCC[j-1]
				}
				maxSCC[i] = l
				break
			}
		}
	}
	fmt.Println(maxSCC)
}

func visit(g *graph, n *node, explored map[string]bool, L []string) []string {
	if !explored[n.label] {
		explored[n.label] = true
		for l := range n.neighborSet {
			L = visit(g, g.getNode(l), explored, L)
		}
		L = append(L, n.label)
	}
	return L
}

func assign(g *graph, n *node, assigned map[string]bool, components map[string][]string, componentName string) {
	if !assigned[n.label] {
		assigned[n.label] = true
		components[componentName] = append(components[componentName], n.label)
		for l := range n.neighborSet {
			assign(g, g.getNode(l), assigned, components, componentName)
		}
	}
}

func (g *graph) readFromFile(path string) *graph {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gReversed := graph{}
	gReversed.init()
	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	ssvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range ssvData {
		g.createNode(line[0])
		g.createNode(line[1])
		g.addNeighbor(line[0], line[1], 1)
		gReversed.createNode(line[1])
		gReversed.createNode(line[0])
		gReversed.addNeighbor(line[1], line[0], 1)
	}
	return &gReversed
}

func (g graph) print() {
	var strBuilder strings.Builder
	for _, v := range g.adList {
		strBuilder.WriteString(fmt.Sprintf("%s\t", v.label))
		for l, c := range v.neighborSet {
			strBuilder.WriteString(fmt.Sprintf("%s(%d)\t", l, c))
		}
		s := strBuilder.String()
		fmt.Println(s[:len(s)-1])
		strBuilder.Reset()
	}
}

func (g *graph) init() {
	g.adList = make(map[string]*node)
}

func (g *graph) addNeighbor(nodeLabel, neighboerLabel string, count uint64) bool {
	n := g.getNode(nodeLabel)
	if n == nil {
		return false
	}
	n.addNeighbor(neighboerLabel, count)
	return true
}

func (g *graph) getNode(label string) *node {
	if g.isNode(label) {
		return g.adList[label]
	}
	return nil
}

func (g *graph) createNode(label string) (*node, bool) {
	isNew := false
	if !g.isNode(label) {
		newNode := &node{
			label: label,
		}
		g.addNode(newNode)
		isNew = true
	}
	return g.getNode(label), isNew
}

func (g *graph) addNode(n *node) bool {
	if g.isNode(n.label) {
		return false
	}
	g.adList[n.label] = n
	return true
}

func (g *graph) isNode(label string) bool {
	_, ok := g.adList[label]
	return ok
}

type graph struct {
	adList map[string]*node
	lock   sync.Mutex
}

func (n *node) addNeighbor(label string, count uint64) {
	if n.neighborSet == nil {
		n.neighborSet = make(map[string]uint64)
	}
	n.neighborSet[label] += count
}

type node struct {
	label       string
	neighborSet map[string]uint64
}

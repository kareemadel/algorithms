package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"sync"

	uuid "github.com/satori/go.uuid"
)

func main() {
	g := graph{}
	g.init()
	g.readFromFile("/home/kareemadel/algorithms/Graphs/Karger_min_cut/test.txt")
	fmt.Println(g.findMinCut(0, 0))
}

func (g graph) size() int {
	return len(g.adList)
}

func (g *graph) init() {
	g.adList = make(map[string]*node)
}

func (g graph) findMinCut(trials uint64, maxGoroutines int) uint64 {
	// use n^2*ln(n) trials to get a high probability of success
	n := uint64(g.size())
	if trials == 0 {
		trials = n * n * uint64(math.Log(float64(n)))
	}
	if maxGoroutines < 1 {
		maxGoroutines = 10000
	}
	// increase or decrease the max number of go routines depending on your machine's spec
	AbsMinCut := make(chan uint64)
	minCutChan := make(chan uint64, maxGoroutines)
	guard := make(chan struct{}, maxGoroutines)
	go func(c chan uint64) {
		minCut := ^uint64(0)
		for i := uint64(0); i < trials; i++ {
			temp := <-minCutChan
			if temp < minCut {
				minCut = temp
			}
		}
		c <- minCut
	}(AbsMinCut)
	for i := uint64(0); i < trials; i++ {
		guard <- struct{}{}
		go func() {
			g.findRandCut(minCutChan)
			<-guard
		}()
	}
	minCut := <-AbsMinCut
	return minCut
}

func (g graph) findRandCut(outChan chan uint64) {
	newG := g.deepCopy()
	size := g.size()
	for size > 2 {
		newG.contractRandomEdge()
		size--
	}
	var minCut uint64
	for _, v := range newG.adList {
		for _, c := range v.neighborSet {
			minCut = uint64(c)
			break
		}
		break
	}
	outChan <- minCut
}

func (g *graph) contractRandomEdge() {
	n1, n2 := g.getRandomEdge()
	g.contract(n1, n2)
}

func (g graph) getRandomEdge() (string, string) {
	i := rand.Intn(g.size())
	var n1, n2 string
	for n1 = range g.adList {
		if i == 0 {
			break
		}
		i--
	}
	n := g.getNode(n1)
	n2 = n.getRandromNeighbor()
	return n1, n2
}

func (g graph) print() {
	var strBuilder strings.Builder
	for _, v := range g.adList {
		fmt.Println(v.label)
		for k := range v.labelSet {
			strBuilder.WriteString(fmt.Sprintf("%s\t", k))
		}
		for l, c := range v.neighborSet {
			strBuilder.WriteString(fmt.Sprintf("%s(%d)\t", l, c))
		}
		s := strBuilder.String()
		fmt.Println(s[:len(s)-1])
		strBuilder.Reset()
	}
}

func (g *graph) deepCopy() graph {
	g.lock.Lock()
	defer g.lock.Unlock()

	newG := graph{}
	newG.init()
	for _, v := range g.adList {
		n, _ := newG.createNode(v.label)
		n.importLabels(v)
		n.importNeighbors(v)
	}
	return newG
}

func (g *graph) readFromFile(path string) *graph {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	tsvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var nodeLabel, neighborLabel string
	for _, line := range tsvData {
		nodeLabel = line[0]
		g.createNode(nodeLabel)
		for j := 1; j < len(line)-1; j++ {
			neighborLabel = line[j]
			g.createNode(neighborLabel)
			g.addNeighbor(nodeLabel, neighborLabel, 1)
		}
	}
	return g
}

func (g graph) isNode(label string) bool {
	_, ok := g.adList[label]
	return ok
}

func (g graph) getNode(label string) *node {
	n, ok := g.adList[label]
	if ok {
		return n
	}
	return nil
}

func (g *graph) createNode(label string) (*node, bool) {
	isNew := false
	if !g.isNode(label) {
		g.adList[label] = &node{
			label:       label,
			labelSet:    map[string]bool{label: true},
			neighborSet: make(map[string]uint64),
		}
		isNew = true
	}
	return g.adList[label], isNew
}

func (g *graph) addNode(n *node) bool {
	if g.isNode(n.label) {
		return false
	}
	g.adList[n.label] = n
	return true
}

func (g *graph) addNeighbor(nodeLabel, neighboerLabel string, count uint64) bool {
	n := g.getNode(nodeLabel)
	if n == nil {
		return false
	}
	n.addNeighbor(neighboerLabel, count)
	return true
}

func (g *graph) removeNode(label string) bool {
	if !g.isNode(label) {
		return false
	}
	delete(g.adList, label)
	return true
}

func (g *graph) contract(n1, n2 string) *node {
	n, m := g.getNode(n1), g.getNode(n2)
	if n == nil && m == nil {
		return nil
	} else if n == nil {
		return m
	} else if m == nil {
		return n
	}
	m.mergeWith(n)
	g.removeNode(n1)
	g.removeNode(n2)
	for l := range m.neighborSet {
		neighbor := g.getNode(l)
		c1 := neighbor.getNeighbor(n1)
		c2 := neighbor.getNeighbor(n2)
		neighbor.addNeighbor(m.label, c1+c2)
		neighbor.removeNeighbor(n1)
		neighbor.removeNeighbor(n2)
	}
	g.addNode(m)
	return m
}

type graph struct {
	adList map[string]*node
	lock   sync.Mutex
}

func (n *node) mergeWith(m *node) {
	label, _ := uuid.NewV4()
	n.label = label.String()
	n.addLabel(label.String())
	n.importLabels(m)
	n.importNeighbors(m)
	n.removeSelfLoops()
}

func (n *node) removeSelfLoops() {
	for k := range n.neighborSet {
		if n.labelSet[k] {
			delete(n.neighborSet, k)
		}
	}
}

func (n *node) importLabels(m *node) {
	for k := range m.labelSet {
		n.labelSet[k] = true
	}
}

func (n *node) importNeighbors(m *node) {
	for k, v := range m.neighborSet {
		n.neighborSet[k] += v
	}
}

func (n *node) getRandromNeighbor() string {
	i := rand.Intn(len(n.neighborSet))
	var label string
	for label = range n.neighborSet {
		if i == 0 {
			break
		}
		i--
	}
	return label
}

func (n *node) getNeighbor(label string) uint64 {
	return n.neighborSet[label]
}

func (n *node) addNeighbor(label string, count uint64) {
	n.neighborSet[label] += count
}

func (n *node) updateNeighbor(label string, count uint64) {
	n.neighborSet[label] = count
}

func (n *node) removeNeighbor(label string) {
	delete(n.neighborSet, label)
}

func (n *node) addLabel(label string) {
	n.labelSet[label] = true
}

type node struct {
	label       string
	labelSet    map[string]bool
	neighborSet map[string]uint64
}

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	originalGraph := readGraph("/home/kareemadel/algorithms/Graphs/Karger_min_cut/test.txt")

	fmt.Println(findMinCut(originalGraph))
}

func findMinCut(graph map[string]*node) uint64 {
	// use n^2*ln(n) trials to get a high probability of success
	n := uint64(len(graph))
	trials := n * n * uint64(math.Log(float64(n)))
	// increase or decrease the max number of go routines depending on your machine's spec
	maxGoroutines := 10000
	AbsMinCut := make(chan uint64)
	minCutChan := make(chan uint64, maxGoroutines)
	guard := make(chan struct{}, maxGoroutines)
	go func(c chan uint64) {
		minCut := ^uint64(0)
		for i := uint64(0); i < trials; i++ {
			temp := <-minCutChan
			minCut = min(temp, minCut)
		}
		c <- minCut
	}(AbsMinCut)
	for i := uint64(0); i < trials; i++ {
		guard <- struct{}{}
		go func() {
			findAMinCut(graph, minCutChan)
			<-guard
		}()
	}
	minCut := <-AbsMinCut
	return minCut
}

func findAMinCut(originalGraph map[string]*node, outChan chan uint64) {
	graph := deepCopy(originalGraph)
	size := len(graph)
	for size > 2 {
		graph = contractRandomly(graph)
		size--
	}
	var minCut uint64
	for _, v := range graph {
		for _, c := range v.neighborSet {
			minCut = uint64(c)
			break
		}
		break
	}
	outChan <- minCut
}

func deepCopy(originalGraph map[string]*node) map[string]*node {
	graph := make(map[string]*node)
	for k, v := range originalGraph {
		newNode := node{
			label:       v.label,
			labelSet:    make(map[string]bool),
			neighborSet: make(map[string]uint64),
		}
		for k, v := range v.labelSet {
			newNode.labelSet[k] = v
		}
		for k, v := range v.neighborSet {
			newNode.neighborSet[k] = v
		}
		graph[k] = &newNode
	}
	return graph
}

func contractRandomly(graph map[string]*node) map[string]*node {
	n1 := getRandomNode(graph)
	n2 := getRandromNeighbor(n1, graph)
	mergedNode := merge(n1, n2)
	delete(graph, n1.label)
	delete(graph, n2.label)
	for l := range mergedNode.neighborSet {
		c1 := graph[l].neighborSet[n1.label]
		c2 := graph[l].neighborSet[n2.label]
		graph[l].neighborSet[mergedNode.label] = c1 + c2
		delete(graph[l].neighborSet, n1.label)
		delete(graph[l].neighborSet, n2.label)
	}
	graph[mergedNode.label] = mergedNode
	return graph
}

func merge(n1, n2 *node) *node {
	label, _ := uuid.NewV4()
	mergedNode := &node{
		label:       label.String(),
		labelSet:    map[string]bool{label.String(): true},
		neighborSet: make(map[string]uint64),
	}
	addLabels(mergedNode, n1)
	addLabels(mergedNode, n2)
	addNeighbors(mergedNode, n1)
	addNeighbors(mergedNode, n2)
	return mergedNode
}

func addLabels(mergedNode, n *node) {
	for k := range n.labelSet {
		mergedNode.labelSet[k] = true
	}
}

func addNeighbors(mergedNode, n *node) {
	for k, v := range n.neighborSet {
		_, ok := mergedNode.labelSet[k]
		if !ok {
			mergedNode.neighborSet[k] += v
		}
	}
}

func getRandomNode(graph map[string]*node) *node {
	i := rand.Intn(len(graph))
	var k string
	for k = range graph {
		if i == 0 {
			break
		}
		i--
	}
	return graph[k]
}

func getRandromNeighbor(n *node, graph map[string]*node) *node {
	i := rand.Intn(len(n.neighborSet))
	var k string
	for k = range n.neighborSet {
		if i == 0 {
			break
		}
		i--
	}
	if graph[k] == nil {
		fmt.Println(k)
	}
	return graph[k]
}

func printGraph(graph map[string]*node) {
	var strBuilder strings.Builder
	for _, v := range graph {
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

func readGraph(path string) map[string]*node {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	nodeMap := make(map[string]*node)
	data := make([]byte, 10)
	isLabel := true
	var nodePtr *node
	var strBuilder strings.Builder
	var label string
	for count, err := file.Read(data); err != io.EOF; count, err = file.Read(data) {
		for i := 0; i < count; i++ {
			if data[i] != '\t' && data[i] != '\n' {
				strBuilder.WriteByte(data[i])
			} else if isLabel && data[i] == '\t' {
				isLabel = false
				label = strBuilder.String()
				strBuilder.Reset()
				nodePtr = nodeMap[label]
				if nodePtr == nil {
					nodeMap[label] = &node{
						label:       label,
						labelSet:    map[string]bool{label: true},
						neighborSet: make(map[string]uint64),
					}
					nodePtr = nodeMap[label]
				}
			} else if data[i] == '\t' {
				label = strBuilder.String()
				strBuilder.Reset()
				_, ok := nodeMap[label]
				if !ok {
					nodeMap[label] = &node{
						label:       label,
						labelSet:    map[string]bool{label: true},
						neighborSet: make(map[string]uint64),
					}
				}
				nodePtr.neighborSet[label]++
			} else if data[i] == '\n' {
				isLabel = true
			}
		}
	}
	return nodeMap
}

type node struct {
	label       string
	labelSet    map[string]bool
	neighborSet map[string]uint64
}

func min(i, j uint64) uint64 {
	if i < j {
		return i
	}
	return j
}
